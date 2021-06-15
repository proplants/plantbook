// Package repo contains implementations of the repository interfaces
package repo

import (
	"context"
	"net"
	"time"

	"github.com/proplants/plantbook/internal/api/models"
	"github.com/proplants/plantbook/pkg/logging"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/zapadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

const (
	pgMaxConns                 int32         = 8
	pgMinConns                 int32         = 4
	pgHealthCheckPeriod        time.Duration = 1 * time.Minute
	pgMaxConnLifetime          time.Duration = 24 * time.Hour
	pgMaxConnIdleTime          time.Duration = 30 * time.Minute
	pgConnConfigConnectTimeout time.Duration = 1 * time.Second
)

// ErrNotFound user not found.
var ErrNotFound error = errors.New("user not found")

// PG ...
type PG struct {
	db *pgxpool.Pool
}

// NewPG builder fog postgres implementation of the repo interface.
func NewPG(ctx context.Context, url string, debug bool) (*PG, error) {
	logger := logging.FromContext(ctx)
	cfg, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, errors.WithMessage(err, "parse url error")
	}
	// Pool соединений обязательно ограничивать сверху,
	// так как иначе есть потенциальная опасность превысить лимит соединений с базой.
	cfg.MaxConns = pgMaxConns
	cfg.MinConns = pgMinConns

	// HealthCheckPeriod - частота проверки работоспособности
	// соединения с Postgres
	cfg.HealthCheckPeriod = pgHealthCheckPeriod

	// MaxConnLifetime - сколько времени будет жить соединение.
	// Так как большого смысла удалять живые соединения нет,
	// можно устанавливать большие значения
	cfg.MaxConnLifetime = pgMaxConnLifetime

	// MaxConnIdleTime - время жизни неиспользуемого соединения,
	// если запросов не поступало, то соединение закроется.
	cfg.MaxConnIdleTime = pgMaxConnIdleTime

	// ConnectTimeout устанавливает ограничение по времени
	// на весь процесс установки соединения и аутентификации.
	cfg.ConnConfig.ConnectTimeout = pgConnConfigConnectTimeout

	// Лимиты в net.Dialer позволяют достичь предсказуемого
	// поведения в случае обрыва сети.
	cfg.ConnConfig.DialFunc = (&net.Dialer{
		KeepAlive: cfg.HealthCheckPeriod,
		// Timeout на установку соединения гарантирует,
		// что не будет зависаний при попытке установить соединение.
		Timeout: cfg.ConnConfig.ConnectTimeout,
	}).DialContext

	// logger
	if logger != nil && debug {
		cfg.ConnConfig.Logger = zapadapter.NewLogger(logger.Desugar())
	}
	pool, err := pgxpool.ConnectConfig(ctx, cfg)
	if err != nil {
		return nil, errors.WithMessage(err, "connect error")
	}
	return &PG{db: pool}, nil
}

func (pg *PG) StoreUser(ctx context.Context, user *models.User, passwordHash []byte) (*models.User, error) {
	const query string = `insert into public.users 
		(name_user, email_addr, pwd_hash,
		first_name, last_name, phone_number, 
		user_role, description)
		values ($1,$2,$3,$4,$5,$6,$7,$8)
		returning id_user;`

	var uid int64
	err := pg.db.QueryRow(ctx, query, user.Username, user.Email, passwordHash,
		user.FirstName, user.LastName, user.Phone,
		user.UserRole, user.Username+" TODO: привести БД в соответствии с моделью").Scan(&uid)
	if err != nil {
		return nil, errors.WithMessage(err, "insert user failed")
	}
	if uid == 0 {
		return nil, errors.Errorf("insert user failed, empty id")
	}
	user.ID = uid
	// TODO: привести в соответствие с моделью и спекой
	// что за статус, что он означает?
	user.UserStatus = 1

	return user, nil
}

func (pg *PG) FindUserByLogin(ctx context.Context, login string) (*models.User, []byte, error) {
	const query string = `SELECT 
			id_user, name_user, email_addr, 
			pwd_hash, first_name, last_name, 
			phone_number, user_role
		FROM public.users where name_user=$1;`
	var u models.User
	var hash []byte
	err := pg.db.QueryRow(ctx, query, login).Scan(
		&u.ID, &u.Username, &u.Email,
		&hash, &u.FirstName, &u.LastName,
		&u.Phone, &u.UserRole)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil, ErrNotFound
		}
		return nil, nil, errors.WithMessage(err, "fetch user failed")
	}
	return &u, hash, nil
}

// Health checks availability postgres.
func (pg *PG) Health(ctx context.Context) error {
	var tmp string
	return pg.db.QueryRow(ctx, "SELECT version();").Scan(&tmp)
}

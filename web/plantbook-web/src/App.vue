<template>
  <v-app>
    <v-app-bar color="lightgrey accent-4" app>
      <v-toolbar-title @click="logout">
        <v-icon color="green" large>mdi-square-rounded</v-icon>
        <router-link :to="{ name: 'Main' }">
          <strong> Plantbook </strong>
        </router-link>
        <button>
          <v-icon @click="$router.go(-1)" color="back" v-if="isLoggedIn"
            >mdi-keyboard-backspace
          </v-icon>
        </button>
      </v-toolbar-title>

      <v-spacer></v-spacer>

      <strong @click="pageUser" @mouseover="active = true">
        {{ LOGIN }}
      </strong>

      <v-btn
        v-if="!isLoggedIn"
        :to="{ name: 'Login' }"
        left
        bottom
        class="grey"
      >
        <strong> Войти </strong>
      </v-btn>
    </v-app-bar>

    <v-main>
      <v-container>
        <v-card
          v-click-outside="onClickOutside"
          @mouseleave="mouseleave"
          v-if="active"
          elevation="8"
          shaped
          class="window"
        >
          <div>
            <p>Мой сад</p>
            <p>В вазе</p>
          </div>
          <strong v-if="isLoggedIn" @click="logout">Выйти</strong>
        </v-card>
        <router-view />
      </v-container>
    </v-main>

    <v-footer padless>
      <v-col class="text-center" cols="12">
        {{ new Date().getFullYear() }}
        <v-icon color="grey" large>mdi-square-rounded</v-icon>
        <a href="https://github.com/proplants/plantbook" target="_blank">
          <strong>Github</strong></a
        >
      </v-col>
    </v-footer>
  </v-app>
</template>

<script>
import { mapGetters, mapActions } from "vuex";

export default {
  name: "App",

  data: () => ({
    active: false,
  }),
  computed: {
    ...mapGetters({
      IS_LOGGED_IN: "auth/IS_LOGGED_IN",
      LOGIN: "auth/LOGIN",
    }),
    isLoggedIn() {
      return this.IS_LOGGED_IN;
    },
  },
  methods: {
    mouseleave() {
      this.active = false;
    },
    pageUser() {
      this.active = false;
      this.$router.push("/UserPage").catch(() => {});
    },
    onClickOutside() {
      this.active = false;
    },
    ...mapActions({
      LOG_OUT: "auth/LOG_OUT",
    }),

    async logout() {
      await this.LOG_OUT();
      this.active = false;
      this.$router.push("/login");
    },
  },
};
</script>
 
<style lang="scss">
.v-application {
  font-family: "Material Design Icons" !important;
  font-size: 18px;
}

a {
  text-decoration: none;
  color: black !important;
}
a,
strong:hover {
  cursor: pointer;
}
.window {
  width: 231px;
  height: 255px;
  border: 1px solid black;
  border-radius: 30px;
  position: fixed !important;
  // right: 85px;
  right: 55px;
  top: 35px;
  background-color: white;
  z-index: 9;
  padding: 15px;
  flex-direction: column;
  display: flex !important;
  justify-content: space-between;

  strong:hover {
    cursor: pointer;
  }
}
</style>
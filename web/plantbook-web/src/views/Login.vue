<template>
  <div class="container">
    <div class="main" align-self="center" align="center">
      <div class="block">
        <h2>Войти</h2>
        <v-alert v-if="showError" color="red" elevation="5" dense type="error">
          Неверный пароль или логин</v-alert
        >
        <form class="login" @submit.prevent="login">
          <v-text-field
            type="name"
            v-model="form.login"
            label="Name"
            required
          ></v-text-field>
          <v-text-field
            type="password"
            v-model="form.password"
            label="Password"
            required
          ></v-text-field>

          <v-btn block class="mr-4" @click="submit"> Войти </v-btn>
        </form>
      </div>
    </div>
  </div>
</template>


<script>
import { mapActions, mapGetters } from "vuex";
export default {
  data: () => ({
    form: {
      login: "",
      password: "",
    },
    showError: false,
  }),
  computed: {
    ...mapGetters({
      IS_LOGGED_IN: "auth/IS_LOGGED_IN",
    }),
  },
  methods: {
    ...mapActions({
      LOGIN: "auth/LOGIN",
    }),
    async submit() {
      await this.LOGIN(this.form);
      this.checkUser();
    },
    checkUser() {
      return this.IS_LOGGED_IN
        ? this.$router.push("/UserGallery")
        : (this.showError = true);
    },
  },
};
</script>


<style lang="scss" scoped>
.block {
  max-width: 300px;

  h2 {
    text-align: left;
  }
}
</style>

<template>
  <v-app>
    <v-app-bar color="lightgrey accent-4" app>
      <v-toolbar-title @click="logout">
        <v-icon color="green" large>mdi-square-rounded</v-icon>
        <router-link :to="{ name: 'Main' }">
          <strong> Plantbook </strong>
        </router-link>
      </v-toolbar-title>

      <v-spacer></v-spacer>

      <v-btn
        v-if="!isLoggedIn"
        :to="{ name: 'Login' }"
        left
        bottom
        class="grey"
      >
        <strong> Войти </strong>
      </v-btn>
      <v-btn v-else :to="{ name: 'Login' }" @click="logout">
        <strong>Выйти</strong>
      </v-btn>
    </v-app-bar>

    <v-main>
      <v-container fluid>
        <router-view />
      </v-container>
    </v-main>

    <v-footer padless>
      <v-col class="text-center" cols="12">
        <v-icon color="grey" large>mdi-square-rounded</v-icon>
        <!-- {{ new Date().getFullYear() }} — -->
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
    //
  }),
  computed: {
    ...mapGetters(["IS_LOGGED_IN"]),
    isLoggedIn() {
      return this.IS_LOGGED_IN;
    },
  },
  methods: {
    ...mapActions(["LOG_OUT"]),

    async logout() {
      await this.LOG_OUT();
      // this.$router.push("/login");
    },
  },
};
</script>

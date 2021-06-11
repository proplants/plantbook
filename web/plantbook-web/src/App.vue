<template>
  <v-app>
    <v-app-bar color="lightgrey accent-4" app>
      <v-toolbar-title> <strong> Plantbook </strong> </v-toolbar-title>

      <v-spacer></v-spacer>

      <v-btn :to="{ name: 'Login' }" left bottom class="grey">
        <strong>Войти </strong>
      </v-btn>
    </v-app-bar>

    <v-main>
      <v-container fluid>
        <router-view />
      </v-container>
    </v-main>

    <v-footer padless>
      <v-col class="text-center" cols="12">
        {{ new Date().getFullYear() }} — <strong>Github</strong>
      </v-col>
    </v-footer>
  </v-app>
</template>

<script>
export default {
  name: "App",

  data: () => ({
    //
  }),
};
</script>

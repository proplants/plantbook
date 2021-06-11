<template>
  <div class="main" align="center">
    <div class="block">
      <h2>Войти</h2>
      <form class="login" @submit.prevent="login">
        <v-text-field v-model="name" label="Name" required></v-text-field>
        <v-text-field
          v-model="password"
          label="Password"
          required
        ></v-text-field>

        <v-btn block class="mr-4 greenyellow" @click="submit"> Войти </v-btn>
      </form>
    </div>
  </div>
</template>


<script>
import { mapActions } from "vuex";
export default {
  data: () => ({
    name: "",
    password: "",
  }),

  computed: {},

  methods: {
    ...mapActions(["LOGIN"]),
    submit() {
      let login = this.name;
      let password = this.password;
      this.LOGIN({ login, password })
        .then(() => this.$router.push("/UserGallery"))
        .catch((err) => console.log(err));
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

<template>
  <v-hover v-slot="{ hover }" open-delay="200">
    <v-card
      class="mx-auto card"
      max-width="344"
      :elevation="hover ? 16 : 2"
      :class="{ 'on-hover': hover }"
    >
      <v-card-title class="mx-auto text-h4"> Добавить </v-card-title>
      <v-card-title class="mx-auto text-h4"> растение </v-card-title>
    </v-card>
  </v-hover>
</template>

<script>
export default {
  data: () => ({}),
};
</script>

<style lang="scss" scoped>
.v-card.on-hover.theme--dark {
  background-color: rgba(#fff, 0.8);
}
.v-card__title {
  color: #514b44;
  letter-spacing: 9px !important;
  font-weight: 600;
  padding: 3px;
}
.card {
  height: 70%;
  margin: auto;
  //
  background-color: #efe9e940;
  display: flex;
  flex-direction: column;
  justify-content: center;
}
</style>
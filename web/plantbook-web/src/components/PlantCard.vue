<template>
  <v-hover v-slot="{ hover }" open-delay="200">
    <v-card
      class="mx-auto"
      max-width="344"
      :elevation="hover ? 16 : 2"
      :class="{ 'on-hover': hover }"
    >
      <v-img
        src="https://cdn.vuetifyjs.com/images/cards/sunshine.jpg"
        height="200px"
      ></v-img>

      <v-card-title> {{ card_item.title }} </v-card-title>

      <v-card-actions>
        <v-btn color="orange lighten-2" text> Explore </v-btn>

        <v-spacer></v-spacer>

        <v-btn icon @click="show = !show">
          <v-icon>{{ show ? "mdi-chevron-up" : "mdi-chevron-down" }}</v-icon>
        </v-btn>
      </v-card-actions>

      <v-expand-transition>
        <div v-show="show">
          <v-divider></v-divider>

          <v-card-text>
            {{ card_item.description }}
          </v-card-text>
        </div>
      </v-expand-transition>
    </v-card>
  </v-hover>
</template>

<script>
export default {
  props: {
    card_item: Object,
  },
  data: () => ({
    show: false,
  }),
};
</script>

<style lang="scss" scoped>
.v-card.on-hover.theme--dark {
  background-color: rgba(#fff, 0.8);
}
> .v-card__text {
  color: #000;
}
</style>
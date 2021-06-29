<template>
  <div class="main">
    <div class="search">
      <v-text-field
        v-model="searchLine"
        label="Введите имя"
        hide-details="auto"
      ></v-text-field>
      <p v-if="!filterCards.length" class="mt-5">
        По заданному имени <strong>{{ searchLine }}</strong> ничего не найдено
      </p>
    </div>

    <section class="garden">
      <div v-for="item in filterCards.slice(0, page)" :key="item.id">
        <PlantCard :card_item="item" />
      </div>
    </section>
    <v-btn @click="next" :disabled="page >= numPages"> Показать больше </v-btn>
  </div>
</template>

<script>
import PlantCard from "../components/PlantCard.vue";
import { mapActions, mapGetters } from "vuex";
export default {
  name: "UserGallery",
  components: {
    PlantCard,
  },
  data() {
    return {
      page: 3,
      gardens: [],
      searchLine: "",
    };
  },
  computed: {
    ...mapGetters({ GET_GARDENS: "gardens/GET_GARDENS" }),
    numPages() {
      return Math.ceil(this.gardens.length);
    },
    filterCards() {
      let regExp = new RegExp(this.searchLine, "i");
      return this.gardens.filter((el) => regExp.test(el.title));
    },
  },
  methods: {
    next() {
      this.page += 3;
    },
    ...mapActions({
      GARDENS: "gardens/GARDENS",
    }),
  },
  async mounted() {
    await this.GARDENS();
    this.gardens = this.GET_GARDENS;
  },
};
</script>

<style lang="scss" scoped>
.main {
  text-align: center;

  button {
    margin: 27px;
  }
}
.garden {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(344px, 1fr));
  gap: 2em 20px;
}
.search {
  padding-bottom: 40px;
  margin: auto;
}
</style>

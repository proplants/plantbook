<template>
  <div>
    <h1 v-if="plant">{{ plant.title }}</h1>
  </div>
</template>


<script>
import { mapGetters } from "vuex";
export default {
  name: "Plant",
  data() {
    return {
      plant: null,
    };
  },
  computed: {
    ...mapGetters({ GET_GARDENS: "gardens/GET_GARDENS" }),
  },
  created() {
    const plant = this.GET_GARDENS.find(
      (item) => item.title == this.$route.params.title
    );
    this.plant = plant;
  },
};
</script>


<style lang="scss" scoped>
</style>
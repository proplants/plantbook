<template>
  <div class="main">
    <div class="blockUser">
      <div class="avatar"></div>
      <h1>{{ LOGIN }}</h1>
      <router-link :to="{ name: 'UserGallery' }"> back </router-link>
    </div>
    <button :class="{ active: garden }" @click="garden = true">Мой сад</button>
    <button class="btn" :class="{ active: !garden }" @click="garden = false">
      &laquo;В вазе&raquo;
    </button>
    <hr class="line" />

    <div class="myGarden" v-show="garden">
      <PlantCard :card_item="plantCard" />
      <AddPlant class="size" />
    </div>
    <div class="basket" v-show="!garden">
      <p>В вашей вазе пока нет ни одного цветка</p>
    </div>
  </div>
</template>


<script>
import { mapGetters } from "vuex";
import PlantCard from "../components/PlantCard.vue";
import AddPlant from "../components/AddPlant.vue";
export default {
  name: "UserPage",
  components: {
    PlantCard,
    AddPlant,
  },
  data() {
    return {
      garden: true,
      plantCard: {
        id: 0,
        userId: 0,
        title: "agave",
        description: "string",
      },
    };
  },
  computed: {
    ...mapGetters({ LOGIN: "auth/LOGIN" }),
  },
};
</script>

<style lang="scss" scoped>
.size {
  width: 100%;
}

.avatar {
  vertical-align: middle;
  width: 70px;
  height: 70px;
  border-radius: 50%;
  background: #e3d6c4;
}
.line {
  padding: 0;
  height: 10px;
  border: none;
  border-top: 1px solid #333;
  box-shadow: 0 10px 10px -10px #8c8b8b inset;
  margin-bottom: 40px;
}
.myGarden {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(344px, 1fr));
  gap: 2em 20px;
  margin-top: 30px;
}

.active {
  color: black;
  font-weight: 700;
}
button {
  color: #8c8b8b;
}
.btn {
  border-left: 3px solid;
  margin-left: 11px;
  padding-left: 11px;
  margin-bottom: 3px;
}
.blockUser {
  padding: 20px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}
.basket {
  display: flex;
  justify-content: center;
}
</style>
<template>
  <div>
    <!-- Search bar -->
    <div class="field">
      <div class="control has-icons-left">
        <input
          class="input"
          type="text"
          placeholder="Search items"
          v-model="search"
          v-on:keyup="searchItems"
        >
        <span class="icon is-small is-left">
          <font-awesome-icon icon="search"/>
        </span>
      </div>
    </div>
    <!-- Search results -->
    <div class="buttons" v-if="searchResults !== []">
      <button
        class="button"
        v-for="item in searchResults"
        :key="item"
        :class="[selectedItems.includes(item) ? 'is-link' : '']"
        @click="selectItem(item)"
      >{{ displayItem(item) }}</button>
    </div>
  </div>
</template>

<script>
export default {
  computed: {
    items() {
      return this.$store.state.items;
    }
  },
  data() {
    return {
      search: "",
      searchResults: [],
      selectedItems: []
    };
  },
  methods: {
    searchItems() {
      if (this.search != "") {
        var tmpSearch = this.search.toLowerCase();
        this.searchResults = this.items.filter(item =>
          item["Name"].toLowerCase().includes(tmpSearch)
        );
      } else {
        this.searchResults = [];
      }
    },
    selectItem(item) {
      this.selectedItems.push(item);
      this.$emit("selectItems", this.selectedItems);
      this.$itemEventBus.$emit("selectItemsToBus", this.selectedItems);
    },
    deselectItemsFromBus() {
      this.selectedItems = [];
    },
    selectItemsFromBus(items) {
      this.selectedItems = items;
    }
  },
  created() {
    this.$itemEventBus.$on("selectItemsToBus", this.selectItemsFromBus);
    this.$itemEventBus.$on("deselectItemsToBus", this.deselectItemsFromBus);
  }
};
</script>
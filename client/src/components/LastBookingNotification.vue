<template>
  <div id="lastBookingNotification">
    <button class="button is-link">{{lastEntry}}</button>
  </div>
</template>

<script>
export default {
  data: function() {
    return {
      lastEntry: ""
    };
  },
  methods: {
    getLastBookings: function() {
      //TODO magic number: 1
      this.$http.post("/getLastNBookings", 1).then(response => {
        var entry = response.body;
        console.log(entry);
        // this.lastEntry.push(
        //   // entry.TimeStamp.toString(),
        //   this.displayUserName(this.getUser(entry.UserID)),
        //   this.displayItem(this.getItem(entry.ItemID)),
        //   entry.Amount,
        //   "x"
        // );
        this.lastEntry += `${this.displayUserName(
          this.getUser(entry.UserID)
        )}, ${entry.Amount}x ${this.displayItem(this.getItem(entry.ItemID))}`;
        // `${entry.TimeStamp.()}`
        console.log(this.lastEntry);
      });
    },
    getUser: function(id) {
      return this.users.find(u => {
        return u.UserID == id;
      });
    },
    getItem: function(id) {
      return this.items.find(i => {
        return i.ItemID == id;
      });
    }
  },
  created() {
    this.$nextTick(this.getLastBookings());
  }
};
</script>

<style>
#lastBookingNotification {
  position: absolute;
  bottom: 0;
  left: 0;
}
</style>

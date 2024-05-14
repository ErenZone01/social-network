<template>
 
</template>

<script lang="js" setup>
import Headers from './Header.vue'
import Posts from './Post.vue'
import Notif from './Notification.vue'
import sharedData from '../assets/js/data.js';
</script>
<script lang="js">
import commonMixin from '../assets/js/untils.js';
export default {
  name: 'Group',
  mixins: [commonMixin],
 
  mounted() {
    this.GetData();
  },
};
 // Function to toggle the display of the form

const closeFormEvent = () => {
  showFormEvent.value = false;
  groupName.value = "";
  groupDescription.value = "";
};
</script>

<style scoped>/* Si vous utilisez CSS */
.bg-custom-gray {
  background-color: #d4cdcd;
  font-size: 100%;
  width: max-content;
}</style>
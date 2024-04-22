<template>
  <div class="box p-5 px-6">
    <div class="flex items-ce justify-between text-black dark:text-white">
      <h3 class="font-bold text-lg">
        Friends
        <span
          class="block text-sm text-gray-500 mt-0. font-normal dark:text-white"
        >
          3489 Friends
        </span>
      </h3>
      <a href="#" class="text-sm text-blue-500">Find Friend</a>
    </div>

    <div class="grid grid-cols-3 gap-2 gap-y-5 text-center text-sm mt-4 mb-2">
      <div v-for="following in sharedData.MyFollowings" :key="following.ID">
        <router-link
          v-if="following.ID"
          :to="{ name: 'TimelineProfile', params: { userID: following.ID } }"
        >
          <div
            v-if="following.Avatar != ''"
            class="relative w-full aspect-square rounded-lg overflow-hidden"
          >
            <img
              :src="'/src/assets/images/avatars/' + following.Avatar"
              alt=""
              class="object-cover w-full h-full inset-0"
            />
          </div>
          <div v-if="following.Nickname != ''" class="mt-2 line-clamp-1">
            {{ following.Nickname }}
          </div></router-link
        >
      </div>
    </div>
  </div>
</template>


<script lang="js" setup>import sharedData from '../../assets/js/data.js';
</script>


<script lang="js">
import commonMixin from '../../assets/js/untils.js';
export default {
  name: 'Friend',
  mixins: [commonMixin],
};
</script>
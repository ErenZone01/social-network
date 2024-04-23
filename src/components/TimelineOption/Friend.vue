<template>
  <div class="flex space-x-4">
    <div
    class="flex-1"
      uk-sticky="media: 1024; end: #js-oversized; offset: 80"
    >
      <div class="box p-5 px-6">
        <div
          class="flex items-baseline justify-between text-black dark:text-white"
        >
          <h3 class="font-bold text-base">My Followings</h3>
        </div>

        <div class="side-list">
          <div
          v-for="following in sharedData.MyFollowings" :key="following.ID"
            class="side-list-item"
          >
            <router-link
              v-if="following.ID"
              :to="{ name: 'TimelineProfile', params: { userID: following.ID } }"
              ><a>
                <img
                  :src="'/src/assets/images/avatars/' + following.Avatar"
                  alt=""
                  class="side-list-image rounded-full"
                />
              </a>
            </router-link>

            <div class="flex-1">
              <a href="timeline.html">
                <h4 class="side-list-title">{{ following.Nickname }}</h4>
              </a>
              <!-- <div class="side-list-info">125k Following</div> -->
            </div>
            <button
              @click="UnFollow(following, $event)"
              class="button bg-primary-soft text-primary dark:text-white"
            >
              unfollow
            </button>
          </div>

          <button class="bg-secondery button w-full mt-2 hidden">
            See all
          </button>
        </div>
      </div>
    </div>
    <div
    class="flex-1"
      uk-sticky="media: 1024; end: #js-oversized; offset: 80"
    >
      <div class="box p-5 px-6">
        <div
          class="flex items-baseline justify-between text-black dark:text-white"
        >
          <h3 class="font-bold text-base">My Followers</h3>
        </div>

        <div class="side-list">
          <div
          v-for="follower in sharedData.MyFollowers" :key="follower.ID"
            class="side-list-item"
          >
            <router-link
              v-if="follower.ID"
              :to="{ name: 'TimelineProfile', params: { userID: follower.ID } }"
              ><a>
                <img
                  :src="'/src/assets/images/avatars/' + follower.Avatar"
                  alt=""
                  class="side-list-image rounded-full"
                />
              </a>
            </router-link>

            <div class="flex-1">
              <a href="timeline.html">
                <h4 class="side-list-title">{{ follower.Nickname }}</h4>
              </a>
              <!-- <div class="side-list-info">125k follower</div> -->
            </div>
            <button
              @click="Follow(follower, $event)"
              class="button bg-primary-soft text-primary dark:text-white"
            >
              follow
            </button>
          </div>

          <button class="bg-secondery button w-full mt-2 hidden">
            See all
          </button>
        </div>
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
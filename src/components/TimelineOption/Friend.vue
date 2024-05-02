<template>
  <div v-if="sharedData.MyuserProfile.Privacy == 'public' || sharedData.Myaccount.ID == sharedData.MyuserProfile.ID || IsMyAccountFollowed()" class="flex space-x-4">
    <div class="flex-1" uk-sticky="media: 1024; end: #js-oversized; offset: 80">
      <div class="box p-5 px-6">
        <div
          class="flex items-baseline justify-between text-black dark:text-white"
        >
          <h3 class="font-bold text-base">Followings</h3>
        </div>

        <div class="side-list">
          <div
            v-for="following in sharedData.MyProfileFollowings"
            :key="following.ID"
            class="side-list-item"
          >
            <router-link
              v-if="following.ID"
              :to="{
                name: 'TimelineProfile',
                params: { userID: following.ID },
              }"
              ><a>
                <img
                  v-if="following.Avatar != ''"
                  :src="'/src/assets/images/avatars/' + following.Avatar"
                  alt=""
                  class="side-list-image rounded-full"
                />
                <img
                  v-else
                  :src="'/src/assets/images/avatars/Avatar.webp'"
                  alt=""
                  class="side-list-image rounded-full"
                />
              </a>
            </router-link>

            <div class="flex-1">
              <a>
                <h4 class="side-list-title">{{ following.Nickname }}</h4>
              </a>
              <!-- <div class="side-list-info">125k Following</div> -->
            </div>
            <button v-if="sharedData.Myaccount.ID == sharedData.MyuserProfile.ID "
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
    <div class="flex-1" uk-sticky="media: 1024; end: #js-oversized; offset: 80">
      <div class="box p-5 px-6">
        <div
          class="flex items-baseline justify-between text-black dark:text-white"
        >
          <h3 class="font-bold text-base">Followers</h3>
        </div>

        <div class="side-list">
          <div
            v-for="follower in sharedData.MyProfileFollowers"
            :key="follower.ID"
            class="side-list-item"
          >
            <router-link
              v-if="follower.ID"
              :to="{ name: 'TimelineProfile', params: { userID: follower.ID } }"
              ><a>
                <img
                  v-if="follower.Avatar != ''"
                  :src="'/src/assets/images/avatars/' + follower.Avatar"
                  alt=""
                  class="side-list-image rounded-full"
                />
                <img
                  v-else
                  :src="'/src/assets/images/avatars/Avatar.webp'"
                  alt=""
                  class="side-list-image rounded-full"
                />
              </a>
            </router-link>

            <div class="flex-1">
              <a>
                <h4 class="side-list-title">{{ follower.Nickname }}</h4>
              </a>
              <!-- <div class="side-list-info">125k follower</div> -->
            </div>
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
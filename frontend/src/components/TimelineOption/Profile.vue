<template>
  <div
    v-if="
      sharedData.MyuserProfile.Privacy == 'public' ||
      sharedData.Myaccount.ID == sharedData.MyuserProfile.ID ||
      IsMyAccountFollowed()
    "
    class="flex 2xl:gap-12 gap-10 mt-8 max-lg:flex-col"
    id="js-oversized"
  >
    <!-- feed story -->

    <div class="flex-1 xl:space-y-6 space-y-3">
      <div
        v-for="Post in sharedData.MyProfilePost"
        :key="Post.Id"
        class="bg-white rounded-xl shadow-sm text-sm font-medium border1 dark:bg-dark2"
      >
        <!-- post heading -->
        <div class="flex gap-3 sm:p-4 p-2.5 text-sm font-medium">
          <a>
            <router-link
              v-if="sharedData.MyuserProfile.ID"
              :to="{
                name: 'TimelineProfile',
                params: { userID: sharedData.MyuserProfile.ID },
              }"
            >
              <img
                v-if="
                  sharedData.MyuserProfile.ID == Post.ID_User &&
                  sharedData.MyuserProfile.Avatar != ''
                "
                :src="
                  '/src/assets/images/avatars/' +
                  sharedData.MyuserProfile.Avatar
                "
                alt=""
                class="w-9 h-9 rounded-full"
              />
              <img
                v-else-if="
                  sharedData.MyuserProfile.ID == Post.ID_User &&
                  sharedData.MyuserProfile.Avatar == ''
                "
                :src="'/src/assets/images/avatars/Avatar.webp'"
                alt=""
                class="w-9 h-9 rounded-full"
              />
            </router-link>
          </a>
          <div class="flex-1">
            <a>
              <h4 class="text-black dark:text-white">
                {{ Post.Names }}
              </h4>
            </a>
            <div class="text-xs text-gray-500 dark:text-white/80">
              {{ Post.CreatedPost }}
            </div>
          </div>
        </div>

        <div class="sm:px-4 p-2.5 pt-0">
          <p>{{ Post.Content }}</p>
        </div>
        <a
          v-if="Post.Image != ''"
          href="#preview_modal"
          uk-toggle=""
          aria-expanded="false"
        >
          <div class="relative w-full lg:h-96 h-full sm:px-4">
            <img
              :src="'/src/assets/images/post/' + Post.Image"
              alt=""
              class="sm:rounded-lg w-full h-full object-cover"
            />
          </div>
        </a>
        <!-- post icons -->
        <div
          class="sm:p-4 p-2.5 flex items-center gap-4 text-xs font-semibold"
        ></div>

        <!-- comments -->
        <div v-for="Comment in sharedData.Allcomment" :key="Comment.Id">
          <div
            class="sm:p-4 p-2.5 border-t border-gray-100 font-normal space-y-3 relative dark:border-slate-700/40"
            v-if="Comment.ID_Post == Post.ID"
          >
            <div
              class="flex items-start gap-3 relative"
              style="background-color: rgb(234 238 243)"
            >
              <a>
                <img
                  v-if="
                    sharedData.AllUtilisateur[Comment.ID_User - 1].Avatar != ''
                  "
                  :src="
                    '/src/assets/images/avatars/' +
                    sharedData.AllUtilisateur[Comment.ID_User - 1].Avatar
                  "
                  alt=""
                  class="w-6 h-6 mt-1 rounded-full"
                />
                <img
                  v-else-if="
                    sharedData.AllUtilisateur[Comment.ID_User - 1].Avatar == ''
                  "
                  :src="'/src/assets/images/avatars/Avatar.webp'"
                  alt=""
                  class="w-6 h-6 mt-1 rounded-full"
                />
              </a>
              <a class="text-black font-medium inline-block dark:text-white">
                {{ Comment.Names }}
              </a>
              <div v-if="Comment.Image != ''">
                <img
                  :src="'/src/assets/images/comment/' + Comment.Image"
                  alt=""
                  class="sm:rounded-lg w-full h-full object-cover"
                />
              </div>
              <div>
                <p class="mt-0.5">{{ Comment.Content }}</p>
              </div>
            </div>
          </div>
        </div>
        <button
          type="button"
          class="flex items-center gap-1.5 text-gray-500 hover:text-blue-500 mt-2"
        >
          <ion-icon
            name="chevron-down-outline"
            class="ml-auto duration-200 group-aria-expanded:rotate-180"
          ></ion-icon>
          More Comment
        </button>
      </div>
    </div>

    <!-- sidebar -->

    <div class="lg:w-[400px]">
      <div
        class="lg:space-y-4 lg:pb-8 max-lg:grid sm:grid-cols-2 max-lg:gap-6"
        uk-sticky="media: 1024; end: #js-oversized; offset: 80"
      >
        <div class="box p-5 px-6">
          <div class="flex items-ce justify-between text-black dark:text-white">
            <h3 class="font-bold text-lg">Info</h3>
            <a href="#" class="text-sm text-blue-500">Edit</a>
          </div>

          <ul class="text-gray-700 space-y-4 mt-4 text-sm dark:text-white/80">
            <li class="flex items-center gap-3">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke-width="1.5"
                stroke="currentColor"
                class="w-6 h-6"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  d="M12.75 19.5v-.75a7.5 7.5 0 00-7.5-7.5H4.5m0-6.75h.75c7.87 0 14.25 6.38 14.25 14.25v.75M6 18.75a.75.75 0 11-1.5 0 .75.75 0 011.5 0z"
                />
              </svg>
              <div>
                Flowwed By
                <span
                  v-if="sharedData.MyProfileFollowers"
                  class="font-semibold text-black dark:text-white"
                >
                  {{ sharedData.MyProfileFollowers.length }} People(s)
                </span>
                <span
                  v-else-if="!sharedData.MyProfileFollowers"
                  class="font-semibold text-black dark:text-white"
                >
                  0 People(s)
                </span>
              </div>
            </li>
            <li class="flex items-center gap-3">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke-width="1.5"
                stroke="currentColor"
                class="w-6 h-6"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  d="M20.94 7.29l-8-6A1 1 0 0012 1a1 1 0 00-.94.71l-8 6A1 1 0 003 8v12a1 1 0 001 1h16a1 1 0 001-1V8a1 1 0 00-.06-.36zM12 14l-8-5h16l-8 5z"
                />
              </svg>
              <div>
                Email
                <span class="font-semibold text-black dark:text-white">
                  {{ sharedData.MyuserProfile.Email }}
                </span>
              </div>
            </li>
            <li class="flex items-center gap-3">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 24 24"
                class="h-6 w-6"
              >
                <circle
                  cx="12"
                  cy="8"
                  r="5"
                  fill="none"
                  stroke="currentColor"
                  stroke-width="1.5"
                />
                <path
                  d="M3 20v-2a4 4 0 014-4h10a4 4 0 014 4v2H3z"
                  fill="none"
                  stroke="currentColor"
                  stroke-width="1.5"
                />
              </svg>

              <div>
                Username
                <span class="font-semibold text-black dark:text-white">
                  {{ sharedData.MyuserProfile.Nickname }}
                </span>
              </div>
            </li>
            <li class="flex items-center gap-3">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
                class="w-6 h-6"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="1.5"
                  d="M12 1c-3.31 0-6 2.69-6 6 0 1.66.68 3.17 1.76 4.24C6.83 12.8 12 19 12 19s5.17-6.2 4.24-7.76A5.96 5.96 0 0 1 18 7c0-3.31-2.69-6-6-6zM9 16c-.83 0-1.5-.67-1.5-1.5S8.17 13 9 13s1.5.67 1.5 1.5S9.83 16 9 16zm6 0c-.83 0-1.5-.67-1.5-1.5S14.17 13 15 13s1.5.67 1.5 1.5S15.83 16 15 16zm-1-6v-2m-2 2v-2m4 2v-2"
                />
              </svg>

              <div>
                Birth
                <span class="font-semibold text-black dark:text-white">
                  {{ sharedData.MyuserProfile.Birth }}
                </span>
              </div>
            </li>
            <li class="flex items-center gap-3">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
                class="w-6 h-6"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="1.5"
                  d="M12 2a5 5 0 110 10 5 5 0 010-10zm0 0v8m0 0v4m0-4h4m-4 0H8m10 0a5 5 0 00-4 8 5 5 0 004-8zm-6 10v2"
                />
              </svg>

              <div>
                About
                <span class="font-semibold text-black dark:text-white">
                  {{ sharedData.MyuserProfile.About }}
                </span>
              </div>
            </li>
          </ul>
        </div>
      </div>
    </div>
  </div>
  <Comment></Comment></template>
<script lang="js" setup>import sharedData from '../../assets/js/data.js'; import Comment from "../Comment.vue";
</script>

<script lang="js">
import commonMixin from '../../assets/js/untils.js';
export default {
  name: 'Friend',
  mixins: [commonMixin],
  async mounted() {
      await this.FetchCustomRef();
  }
};
 </script>


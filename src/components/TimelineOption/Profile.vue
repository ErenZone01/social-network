<template>
  <div class="flex 2xl:gap-12 gap-10 mt-8 max-lg:flex-col" id="js-oversized">
    <!-- feed story -->

    <div class="flex-1 xl:space-y-6 space-y-3">
      <!--  post image-->
      <div
        class="bg-white rounded-xl shadow-sm text-sm font-medium border1 dark:bg-dark2"
      >
        <!-- post heading -->
        <div class="flex gap-3 sm:p-4 p-2.5 text-sm font-medium">
          <a href="timeline.html">
            <img
              src="/src/assets/images/avatars/avatar-3.jpg"
              alt=""
              class="w-9 h-9 rounded-full"
            />
          </a>
          <div class="flex-1">
            <a href="timeline.html">
              <h4 class="text-black dark:text-white">Monroe Parker</h4>
            </a>
            <div class="text-xs text-gray-500 dark:text-white/80">
              2 hours ago
            </div>
          </div>

          <div class="-mr-1">
            <button type="button" class="button-icon w-8 h-8">
              <ion-icon class="text-xl" name="ellipsis-horizontal"></ion-icon>
            </button>
            <div
              class="w-[245px]"
              uk-dropdown="pos: bottom-right; animation: uk-animation-scale-up uk-transform-origin-top-right; animate-out: true; mode: click"
            >
              <nav>
                <a href="#">
                  <ion-icon
                    class="text-xl shrink-0"
                    name="bookmark-outline"
                  ></ion-icon>
                  Add to favorites
                </a>
                <a href="#">
                  <ion-icon
                    class="text-xl shrink-0"
                    name="notifications-off-outline"
                  ></ion-icon>
                  Mute Notification
                </a>
                <a href="#">
                  <ion-icon
                    class="text-xl shrink-0"
                    name="flag-outline"
                  ></ion-icon>
                  Report this post
                </a>
                <a href="#">
                  <ion-icon
                    class="text-xl shrink-0"
                    name="share-outline"
                  ></ion-icon>
                  Share your profile
                </a>
                <hr />
                <a
                  href="#"
                  class="text-red-400 hover:!bg-red-50 dark:hover:!bg-red-500/50"
                >
                  <ion-icon
                    class="text-xl shrink-0"
                    name="stop-circle-outline"
                  ></ion-icon>
                  Unfollow
                </a>
              </nav>
            </div>
          </div>
        </div>

        <!-- post image -->
        <div class="relative w-full lg:h-96 h-full sm:px-4">
          <img
            src="/src/assets/images/post/img-2.jpg"
            alt=""
            class="sm:rounded-lg w-full h-full object-cover"
          />
        </div>

        <!-- post icons -->
        <div class="sm:p-4 p-2.5 flex items-center gap-4 text-xs font-semibold">
          <div>
            <div class="flex items-center gap-2.5">
              <button
                type="button"
                class="button-icon text-red-500 bg-red-100 dark:bg-slate-700"
              >
                <ion-icon class="text-lg" name="heart"></ion-icon>
              </button>
              <a href="#">1,300</a>
            </div>
            <div
              class="p-1 px-2 bg-white rounded-full drop-shadow-md w-[212px] dark:bg-slate-700 text-2xl"
              uk-drop="offset:10;pos: top-left; animate-out: true; animation: uk-animation-scale-up uk-transform-origin-bottom-left"
            >
              <div
                class="flex gap-2"
                uk-scrollspy="target: > button; cls: uk-animation-scale-up; delay: 100 ;repeat: true"
              >
                <button
                  type="button"
                  class="text-red-600 hover:scale-125 duration-300"
                >
                  <span> 👍 </span>
                </button>
                <button
                  type="button"
                  class="text-red-600 hover:scale-125 duration-300"
                >
                  <span> ❤️ </span>
                </button>
                <button
                  type="button"
                  class="text-red-600 hover:scale-125 duration-300"
                >
                  <span> 😂 </span>
                </button>
                <button
                  type="button"
                  class="text-red-600 hover:scale-125 duration-300"
                >
                  <span> 😯 </span>
                </button>
                <button
                  type="button"
                  class="text-red-600 hover:scale-125 duration-300"
                >
                  <span> 😢 </span>
                </button>
              </div>

              <div
                class="w-2.5 h-2.5 absolute -bottom-1 left-3 bg-white rotate-45 hidden"
              ></div>
            </div>
          </div>
          <div class="flex items-center gap-3">
            <button
              type="button"
              class="button-icon bg-slate-200/70 dark:bg-slate-700"
            >
              <ion-icon class="text-lg" name="chatbubble-ellipses"></ion-icon>
            </button>
            <span>260</span>
          </div>
          <button type="button" class="button-icon ml-auto">
            <ion-icon class="text-xl" name="paper-plane-outline"></ion-icon>
          </button>
          <button type="button" class="button-icon">
            <ion-icon class="text-xl" name="share-outline"></ion-icon>
          </button>
        </div>

        <!-- comments -->
        <div
          class="sm:p-4 p-2.5 border-t border-gray-100 font-normal space-y-3 relative dark:border-slate-700/40"
        >
          <div class="flex items-start gap-3 relative">
            <a href="timeline.html">
              <img
                src="/src/assets/images/avatars/avatar-2.jpg"
                alt=""
                class="w-6 h-6 mt-1 rounded-full"
              />
            </a>
            <div class="flex-1">
              <a
                href="timeline.html"
                class="text-black font-medium inline-block dark:text-white"
              >
                Steeve
              </a>
              <p class="mt-0.5">What a beautiful photo! I love it. 😍</p>
            </div>
          </div>
          <div class="flex items-start gap-3 relative">
            <a href="timeline.html">
              <img
                src="/src/assets/images/avatars/avatar-3.jpg"
                alt=""
                class="w-6 h-6 mt-1 rounded-full"
              />
            </a>
            <div class="flex-1">
              <a
                href="timeline.html"
                class="text-black font-medium inline-block dark:text-white"
              >
                Monroe
              </a>
              <p class="mt-0.5">You captured the moment.😎</p>
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

        <!-- add comment -->
        <div
          class="sm:px-4 sm:py-3 p-2.5 border-t border-gray-100 flex items-center gap-1 dark:border-slate-700/40"
        >
          <img
            src="/src/assets/images/avatars/avatar-7.jpg"
            alt=""
            class="w-6 h-6 rounded-full"
          />

          <div class="flex-1 relative overflow-hidden h-10">
            <textarea
              placeholder="Add Comment...."
              rows="1"
              class="w-full resize-none !bg-transparent px-4 py-2 focus:!border-transparent focus:!ring-transparent"
            ></textarea>

            <div class="!top-2 pr-2" uk-drop="pos: bottom-right; mode: click">
              <div
                class="flex items-center gap-2"
                uk-scrollspy="target: > svg; cls: uk-animation-slide-right-small; delay: 100 ;repeat: true"
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 24 24"
                  fill="currentColor"
                  class="w-6 h-6 fill-sky-600"
                >
                  <path
                    fill-rule="evenodd"
                    d="M1.5 6a2.25 2.25 0 012.25-2.25h16.5A2.25 2.25 0 0122.5 6v12a2.25 2.25 0 01-2.25 2.25H3.75A2.25 2.25 0 011.5 18V6zM3 16.06V18c0 .414.336.75.75.75h16.5A.75.75 0 0021 18v-1.94l-2.69-2.689a1.5 1.5 0 00-2.12 0l-.88.879.97.97a.75.75 0 11-1.06 1.06l-5.16-5.159a1.5 1.5 0 00-2.12 0L3 16.061zm10.125-7.81a1.125 1.125 0 112.25 0 1.125 1.125 0 01-2.25 0z"
                    clip-rule="evenodd"
                  />
                </svg>
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 20 20"
                  fill="currentColor"
                  class="w-5 h-5 fill-pink-600"
                >
                  <path
                    d="M3.25 4A2.25 2.25 0 001 6.25v7.5A2.25 2.25 0 003.25 16h7.5A2.25 2.25 0 0013 13.75v-7.5A2.25 2.25 0 0010.75 4h-7.5zM19 4.75a.75.75 0 00-1.28-.53l-3 3a.75.75 0 00-.22.53v4.5c0 .199.079.39.22.53l3 3a.75.75 0 001.28-.53V4.75z"
                  />
                </svg>
              </div>
            </div>
          </div>

          <button
            type="submit"
            class="text-sm rounded-full py-1.5 px-3.5 bg-secondery"
          >
            Replay
          </button>
        </div>
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
                  v-if="sharedData.MyFollowers"
                  class="font-semibold text-black dark:text-white"
                >
                  {{ sharedData.MyFollowers.length }} People(s)
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
                  {{ sharedData.Myuser.Email }}
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
                  {{ sharedData.Myuser.Nickname }}
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
                  {{ sharedData.Myuser.Birth }}
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
                  {{ sharedData.Myuser.About }}
                </span>
              </div>
            </li>
          </ul>
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
  async mounted() {
        await this.FetchCustomRef();
  }
};
 </script>


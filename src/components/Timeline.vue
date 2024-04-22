<template>
  <Headers />
  <div id="wrapper">
    <!-- main contents -->
    <main
      id="site__main"
      class="2xl:ml-[--w-side] xl:ml-[--w-side-sm] p-2.5 h-[calc(100vh-var(--m-top))] mt-[--m-top]"
    >
      <div class="max-w-[1065px] mx-auto max-lg:-m-2.5">
        <!-- cover  -->
        <div class="bg-white shadow lg:rounded-b-2xl lg:-mt-10 dark:bg-dark2">
          <!-- cover -->
          <div class="relative overflow-hidden w-full lg:h-72 h-48">
            <img
              v-if="sharedData.Myuser && sharedData.Myuser.Avatar"
              :src="'/src/assets/images/avatars/' + sharedData.Myuser.Avatar"
              alt=""
              class="h-full w-full object-cover inset-0"
            />

            <!-- overly -->
            <div
              class="w-full bottom-0 absolute left-0 bg-gradient-to-t from-black/60 pt-20 z-10"
            ></div>

            <div class="absolute bottom-0 right-0 m-4 z-20">
              <div class="flex items-center gap-3">
                <button
                  class="button bg-black/10 text-white flex items-center gap-2 backdrop-blur-small"
                >
                  Edit
                </button>
              </div>
            </div>
          </div>

          <!-- user info -->
          <div class="p-3">
            <div
              class="flex flex-col justify-center md:items-center lg:-mt-48 -mt-28"
            >
              <div class="relative lg:h-48 lg:w-48 w-28 h-28 mb-4 z-10">
                <div
                  class="relative overflow-hidden rounded-full md:border-[6px] border-gray-100 shrink-0 dark:border-slate-900 shadow"
                >
                  <img
                    :src="'/src/assets/images/avatars/' + sharedData.Myuser.Avatar"
                    alt=""
                    class="h-full w-full object-cover inset-0"
                  />
                </div>
                <button
                  type="button"
                  class="absolute -bottom-3 left-1/2 -translate-x-1/2 bg-white shadow p-1.5 rounded-full sm:flex hidden"
                >
                  <ion-icon
                    name="camera"
                    class="text-2xl md hydrated"
                    role="img"
                    aria-label="camera"
                  ></ion-icon>
                </button>
              </div>

              <h3
                class="md:text-3xl text-base font-bold text-black dark:text-white"
              >
                {{ sharedData.Myuser.Firstname }} {{ sharedData.Myuser.Lastname }}
              </h3>

              <p
                class="mt-2 max-w-xl text-sm md:font-normal font-light text-center hidden"
              >
                I love beauty and emotion. 🥰 I’m passionate about photography
                and learning. 📚 I explore genres and styles. 🌈 I think
                photography is storytelling. 😊
              </p>
            </div>
          </div>

          <!-- navigations -->
          <div
            class="flex items-center justify-between mt-3 border-t border-gray-100 px-2 max-lg:flex-col dark:border-slate-700"
            uk-sticky="offset:50; cls-active: bg-white/80 shadow rounded-b-2xl z-50 backdrop-blur-xl dark:!bg-slate-700/80; animation:uk-animation-slide-top ; media: 992"
          >
            <div
              class="flex items-center gap-2 text-sm py-2 pr-1 max-md:w-full lg:order-2"
            >
              <div>
                <div
                  class="w-[240px]"
                  uk-dropdown="pos: bottom-right; animation: uk-animation-scale-up uk-transform-origin-top-right; animate-out: true; mode: click;offset:10"
                >
                  <nav>
                    <a href="#">
                      <ion-icon
                        class="text-xl"
                        name="pricetags-outline"
                      ></ion-icon>
                      Unfollow
                    </a>
                    <a href="#">
                      <ion-icon class="text-xl" name="time-outline"></ion-icon>
                      Mute story
                    </a>
                    <a href="#">
                      <ion-icon class="text-xl" name="flag-outline"></ion-icon>
                      Report
                    </a>
                    <a href="#">
                      <ion-icon class="text-xl" name="share-outline"></ion-icon>
                      Share profile
                    </a>
                    <hr />
                    <a
                      href="#"
                      class="text-red-400 hover:!bg-red-50 dark:hover:!bg-red-500/50"
                    >
                      <ion-icon
                        class="text-xl"
                        name="stop-circle-outline"
                      ></ion-icon>
                      Block
                    </a>
                  </nav>
                </div>
              </div>
            </div>

            <nav
              class="flex gap-0.5 rounded-xl -mb-px text-gray-600 font-medium text-[15px] dark:text-white max-md:w-full max-md:overflow-x-auto"
            >
              <!-- Lien pour la timeline -->
              <router-link
                :to="{
                  name: 'TimelineProfile',
                  params: { userID: $route.params.userID },
                }"
                class="inline-block py-3 leading-8 px-3.5 border-b-2 border-blue-600 text-blue-600"
                >Timeline</router-link
              >

              <!-- Lien pour les amis -->
              <router-link
                :to="{
                  name: 'TimelineFriends',
                  params: { userID: $route.params.userID },
                }"
                class="inline-block py-3 leading-8 px-3.5"
                >Friend
                <span class="text-xs pl-2 font-normal lg:inline-block hidden"
                  >2,680</span
                ></router-link
              >

              <!-- Lien pour les groupes -->
              <router-link
                :to="{
                  name: 'TimelineGroups',
                  params: { userID: $route.params.userID },
                }"
                class="inline-block py-3 leading-8 px-3.5"
                >Group</router-link
              >
            </nav>
          </div>
        </div>
        <router-view></router-view>
      </div>
    </main>
  </div>
  <Chat></Chat>
  <Posts></Posts>
  <Notif></Notif>
</template>


<script lang="js" setup>
import Headers from "./Header.vue";
import Posts from './Post.vue'
import Chat from './Chat.vue'
import Notif from './Notification.vue'
import { RouterLink } from "vue-router";
import sharedData from '../assets/js/data.js';
</script>

<script lang="js">import commonMixin from '../assets/js/untils.js';
export default {
  name: 'Friend',
  mixins: [commonMixin],
};</script>
<template>
  <Headers />
  <div id="wrapper">
    <!-- main contents -->
    <main
      id="site__main"
      class="2xl:ml-[--w-side] xl:ml-[--w-side-sm] p-2.5 h-[calc(100vh-var(--m-top))] mt-[--m-top]"
    >
      <!-- timeline -->
      <div
        class="lg:flex 2xl:gap-16 gap-12 max-w-[1065px] mx-auto"
        id="js-oversized"
      >
        <div class="max-w-[680px] mx-auto">
          <!-- feed story -->
          <div class="md:max-w-[580px] mx-auto flex-1 xl:space-y-6 space-y-3">
            <!-- add story -->
            <div
              class="bg-white rounded-xl shadow-sm md:p-4 p-2 space-y-4 text-sm font-medium border1 dark:bg-dark2"
            >
              <div
                class="bg-white rounded-xl shadow-sm md:p-4 p-2 space-y-4 text-sm font-medium border1 dark:bg-dark2"
              >
                <div class="flex items-center md:gap-3 gap-1">
                  <div
                    class="flex-1 bg-slate-100 hover:bg-opacity-80 transition-all rounded-lg cursor-pointer dark:bg-dark3"
                    uk-toggle="target: #create-status"
                  >
                    <div class="py-2.5 text-center dark:text-white">
                      create your post here
                    </div>
                  </div>
                </div>
              </div>

              <!--  post image with slider-->

              <!-- post text-->
              <div
                v-for="Post in sharedData.Allpost"
                :key="Post.Id"
                class="bg-white rounded-xl shadow-sm text-sm font-medium border1 dark:bg-dark2"
              >
                <div v-for="user in sharedData.AllUtilisateur" :key="user.ID">
                  <div
                    v-if="
                      (user.ID == Post.ID_User &&
                        (Post.Privacy == 'Public' ||
                          (Post.Privacy == 'Private' &&
                            IfIFollowAnAccount(user)))) ||
                      Post.ID_User == sharedData.Myaccount.ID
                    "
                  >
                    <!-- post heading -->
                    <div class="flex gap-3 sm:p-4 p-2.5 text-sm font-medium">
                      <a>
                        <router-link
                          v-if="user.ID"
                          :to="{
                            name: 'TimelineProfile',
                            params: { userID: user.ID },
                          }"
                        >
                          <img
                            v-if="user.ID == Post.ID_User && user.Avatar != ''"
                            :src="'/src/assets/images/avatars/' + user.Avatar"
                            alt=""
                            class="w-9 h-9 rounded-full"
                          />
                          <img
                            v-else-if="
                              user.ID == Post.ID_User && user.Avatar == ''
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
                    <div
                      v-for="Comment in sharedData.Allcomment"
                      :key="Comment.Id"
                      class="sm:p-4 p-2.5 border-t border-gray-100 font-normal space-y-3 relative dark:border-slate-700/40">
                    <div v-if="Comment.ID_Post == Post.ID">
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
                        
                          </a>
                          <p class="mt-0.5">
                          </p>
                          <img
                          :src="'/src/assets/images/comment/' + Comment.Image"
                          alt=""
                          class="sm:rounded-lg w-full h-full object-cover"
                        />
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
                          {{ Comment.Names }}
                          </a>
                          <p class="mt-0.5">{{ Comment.Content }}</p>
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
                        More Commen:t
                      </button>
                    </div>

                    <!-- add comment -->
               

                      <button
                        type="submit"
                        class="text-sm rounded-full py-1.5 px-3.5 bg-secondery"
                      >
                      </button>
                    </div>
                    <div
                      class="sm:px-4 sm:py-3 p-2.5 border-t border-gray-100 flex items-center gap-1 dark:border-slate-700/40"
                    >
                      <img
                        v-if="user.Avatar != ''"
                        :src="'/src/assets/images/avatars/' + user.Avatar"
                        alt=""
                        class="w-9 h-9 rounded-full"
                      />
                      <img
                        v-else-if="user.Avatar == ''"
                        :src="'/src/assets/images/avatars/Avatar.webp'"
                        alt=""
                        class="w-9 h-9 rounded-full"
                      />

                      <div
                    class="flex-1 bg-slate-100 hover:bg-opacity-80 transition-all rounded-lg cursor-pointer dark:bg-dark3"
                    uk-toggle="target: #create-comment"
                  >
                    <div @click="  this.RecupId(Post.ID)"  class="py-2.5 text-center dark:text-white">
                      add comment
                    </div>
                  </div>
                      </div>
                  </div>
                </div>
              </div>

              <!-- placeholder -->
            </div>
          </div>
        </div>
        <!-- sidebar -->
        <div class="flex-1">
          <div
            class="lg:space-y-4 lg:pb-8 max-lg:grid sm:grid-cols-2 max-lg:gap-6"
            uk-sticky="media: 1024; end: #js-oversized; offset: 80"
          >
            <div class="box p-5 px-6">
              <div
                class="flex items-baseline justify-between text-black dark:text-white"
              >
                <h3 class="font-bold text-base">People you may know</h3>
              </div>
              <div class="side-list">
                <div
                  v-for="user in sharedData.AllUsers"
                  :key="user.ID"
                  class="side-list-item"
                >
                  <router-link
                    v-if="user.ID"
                    :to="{
                      name: 'TimelineProfile',
                      params: { userID: user.ID },
                    }"
                    ><a>
                      <img
                        v-if="user.Avatar != ''"
                        :src="'/src/assets/images/avatars/' + user.Avatar"
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
                      <h4 class="side-list-title">
                        {{ user.Nickname }}
                      </h4>
                    </a>
                    <!-- <div class="side-list-info">125k user</div> -->
                  </div>

                  <button
                    @click="Follow(user, $event)"
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
      </div>
    </main>
  </div>

  <Chat></Chat>
  <Posts></Posts>
  <Notif></Notif>
  <Comment></Comment>
</template>

<script lang="js" setup>
import Headers from "./Header.vue";
import Posts from "./Post.vue";
import Chat from "./Chat.vue";
import Notif from "./Notification.vue";
import Comment from "./Comment.vue";

import sharedData from "../assets/js/data.js";
</script>

<script lang="js">
import commonMixin from "../assets/js/untils.js";
export default {
  name: "Home",
  mixins: [commonMixin],
  mounted() {
    this.GetData();
  },
};
</script>
<style scoped></style>

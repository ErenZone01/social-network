<template>
  <Headers />
  <!-- main contents -->
  <main
    id="site__main"
    class="2xl:ml-[--w-side] xl:ml-[--w-side-sm] p-2.5 h-[calc(100vh-var(--m-top))] mt-[--m-top]"
  >
    <div class="max-w-[1065px] mx-auto">
      <!-- cover  -->
      <div class="bg-white shadow lg:rounded-b-2xl lg:-mt-10 dark:bg-dark2">
        <!-- cover -->
        <div class="relative overflow-hidden w-full lg:h-72 h-36">
          <img
            src="/src/assets/images/group/group-1.jpg"
            alt=""
            class="h-full w-full object-cover inset-0"
          />

          <!-- overly -->
          <div
            class="w-full bottom-0 absolute left-0 bg-gradient-to-t from -black/60 pt-10 z-10"
          ></div>
        </div>

        <div class="lg:px-10 md:p-5 p-3">
          <div class="flex flex-col justify-center">
            <div class="flex lg:items-center justify-between max-md:flex-col">
              <div class="flex-1">
                <h3
                  class="md:text-2xl text-base font-bold text-black dark:text-white"
                >
                  {{ sharedData.MygroupProfile.GroupName }}
                </h3>
                <p
                  class="font-normal text-gray-500 mt-2 flex gap-2 flex-wrap dark:text-white/80"
                >
                  <span class="max-lg:hidden"> Public group </span>
                </p>
              </div>

              <div>
                <div class="flex items-center gap-2 mt-1">
                  <button
                    class="button bg-primary flex items-center gap-1 text-white py-2 px-3.5 shadow ml-auto"
                  >
                    <ion-icon name="add-outline" class="text-xl"></ion-icon>
                    <span class="text-sm"> Leave </span>
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>

        <div
          class="flex items-center justify-between border-t border-gray-100 px-2 dark:border-slate-700"
        >
          <nav
            class="flex gap-0.5 rounded-xl overflow-hidden -mb-px text-gray-500 font-medium text-sm overflow-x-auto dark:text-white"
          >
            <a
              href="#"
              class="inline-block py-3 leading-8 px-3.5 border-b-2 border-blue-600 text-blue-600"
              >Discussion</a
            >

            <a href="#" class="inline-block py-3 leading-8 px-3.5">Event</a>
            <a href="#" class="inline-block py-3 leading-8 px-3.5">Members</a>
          </nav>
        </div>
      </div>

      <div
        class="flex 2xl:gap-12 gap-10 mt-8 max-lg:flex-col"
        id="js-oversized"
      >
        <!-- feed story -->
        <div class="flex-1 xl:space-y-6 space-y-3">
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
              v-for="Post in sharedData.MyProfilePost"
              :key="Post.Id"
              class="bg-white rounded-xl shadow-sm text-sm font-medium border1 dark:bg-dark2"
            >
              <div v-for="user in sharedData.AllUtilisateur" :key="user.ID">
                <div v-if="user.ID == Post.ID_User && Post.Types == 'Group'">
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
                        <p class="mt-0.5">
                          What a beautiful photo! I love it. 😍
                        </p>
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

                    <div class="flex-1 relative overflow-hidden h-10">
                      <textarea
                        placeholder="Add Comment...."
                        rows="1"
                        class="w-full resize-none !bg-transparent px-4 py-2 focus:!border-transparent focus:!ring-transparent"
                      ></textarea>

                      <div
                        class="!top-2 pr-2"
                        uk-drop="pos: bottom-right; mode: click"
                      >
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
            </div>

            <!-- placeholder -->
          </div>
        </div>

        <div class="lg:w-[400px]">
          <div
            class="lg:space-y-4 lg:pb-8 max-lg:grid sm:grid-cols-2 max-lg:gap-6"
            uk-sticky="media: 1024; end: #js-oversized; offset: 80"
          >
            <!-- group info -->
            <div class="box p-5 px-6">
              <div
                class="flex items-ce justify-between text-black dark:text-white"
              >
                <h3 class="font-bold text-lg">Info</h3>
              </div>

              <ul
                class="text-gray-700 space-y-4 mt-2 mb-1 text-sm dark:text-white"
              >
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
                    Description
                    <span class="font-semibold text-black dark:text-white">{{
                      sharedData.MygroupProfile.GroupDescription
                    }}</span>
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
                      d="M15 19.128a9.38 9.38 0 002.625.372 9.337 9.337 0 004.121-.952 4.125 4.125 0 00-7.533-2.493M15 19.128v-.003c0-1.113-.285-2.16-.786-3.07M15 19.128v.106A12.318 12.318 0 018.624 21c-2.331 0-4.512-.645-6.374-1.766l-.001-.109a6.375 6.375 0 0111.964-3.07M12 6.375a3.375 3.375 0 11-6.75 0 3.375 3.375 0 016.75 0zm8.25 2.25a2.625 2.625 0 11-5.25 0 2.625 2.625 0 015.25 0z"
                    />
                  </svg>
                  <div>
                    Members
                    <span
                      v-if="sharedData.MygroupProfile.IdMember"
                      class="font-semibold text-black dark:text-white"
                    >
                      {{ sharedData.MygroupProfile.IdMember.length }} People
                    </span>
                  </div>
                </li>
              </ul>
            </div>

            <div class="box p-5 px-6">
              <div
                class="flex items-baseline justify-between text-black dark:text-white"
              >
                <h3 class="font-bold text-base">People you may know</h3>
              </div>
              <div class="side-list">
                <div
                  v-for="user in sharedData.AllUtilisateur"
                  :key="user.ID"
                 
                >
                  <div class="side-list-item" v-if="(!sharedData.MygroupProfile.IdMember || !sharedData.MygroupProfile.IdMember.includes(user.ID)) && sharedData.MygroupProfile.IdCreator != user.ID">
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
                      @click="this.InvitationGroup(sharedData.MygroupProfile.ID_Group , user.ID, $event)"
                      class="button bg-primary-soft text-primary dark:text-white"
                    >
                      add
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </main>
  <Chat></Chat>
  <Posts></Posts>
  <Notif></Notif>
</template>
<script lang="js" setup>
import Headers from './Header.vue'
import Posts from './Post.vue'
import Chat from './Chat.vue'
import Notif from './Notification.vue'
import sharedData from '../assets/js/data.js';
</script>
<script lang="js">
import commonMixin from '../assets/js/untils.js';
export default {
  name: 'TimelineGroup',
  mixins: [commonMixin],
 
  mounted() {
    this.FetchCustomRefGroup()
  }
};
</script>
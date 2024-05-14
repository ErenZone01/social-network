<template>
  <div class="flex 2xl:gap-12 gap-10 mt-8 max-lg:flex-col" id="js-oversized">
    <!-- feed story -->
    <div class="flex-1 xl:space-y-6 space-y-3">
      <!-- add story -->
      <div class="bg-white rounded-xl shadow-sm md:p-4 p-2 space-y-4 text-sm font-medium border1 dark:bg-dark2">
        <div class="bg-white rounded-xl shadow-sm md:p-4 p-2 space-y-4 text-sm font-medium border1 dark:bg-dark2">
          <div class="flex items-center md:gap-3 gap-1">
            <div class="flex-1 bg-slate-100 hover:bg-opacity-80 transition-all rounded-lg cursor-pointer dark:bg-dark3"
              uk-toggle="target: #create-status">
              <div class="py-2.5 text-center dark:text-white">
                create your post here
              </div>
            </div>
          </div>
        </div>

        <!--  post image with slider-->

        <!-- post text-->
        <div v-for="Post in sharedData.MyProfilePost" :key="Post.Id"
          class="bg-white rounded-xl shadow-sm text-sm font-medium border1 dark:bg-dark2">
          <div v-for="user in sharedData.AllUtilisateur" :key="user.ID">
            <div v-if="user.ID == Post.ID_User && Post.Types == 'Group'">
              <!-- post heading -->
              <div class="flex gap-3 sm:p-4 p-2.5 text-sm font-medium">
                <a>
                  <router-link v-if="user.ID" :to="{
                    name: 'TimelineProfile',
                    params: { userID: user.ID },
                  }">
                    <img v-if="user.ID == Post.ID_User && user.Avatar != ''"
                      :src="'/src/assets/images/avatars/' + user.Avatar" alt="" class="w-9 h-9 rounded-full" />
                    <img v-else-if="user.ID == Post.ID_User && user.Avatar == ''"
                      :src="'/src/assets/images/avatars/Avatar.webp'" alt="" class="w-9 h-9 rounded-full" />
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
                <p class="texte">{{ Post.Content }}</p>
              </div>
              <a v-if="Post.Image != ''" href="#preview_modal" uk-toggle="" aria-expanded="false">
                <div class="relative w-full lg:h-96 h-full sm:px-4">
                  <img :src="'/src/assets/images/post/' + Post.Image" alt=""
                    class="sm:rounded-lg w-full h-full object-cover" />
                </div>
              </a>
              <!-- post icons -->
              <div class="sm:p-4 p-2.5 flex items-center gap-4 text-xs font-semibold"></div>

              <!-- comments -->
              <div v-for="Comment in sharedData.Allcomment" :key="Comment.Id">
                <div
                  class="sm:p-4 p-2.5 border-t border-gray-100 font-normal space-y-3 relative dark:border-slate-700/40"
                  v-if="Comment.ID_Post == Post.ID">
                  <div class="bg-white rounded-xl shadow-sm text-sm font-medium border1 dark:bg-dark2" style="background-color: rgb(234 238 243)">
                    <a>
                      <img v-if="
                        sharedData.AllUtilisateur[Comment.ID_User - 1]
                          .Avatar != ''
                      " :src="'/src/assets/images/avatars/' +
                          sharedData.AllUtilisateur[Comment.ID_User - 1].Avatar
                          " alt="" class="w-6 h-6 mt-1 rounded-full" />
                      <img v-else-if="
                        sharedData.AllUtilisateur[Comment.ID_User - 1]
                          .Avatar == ''
                      " :src="'/src/assets/images/avatars/Avatar.webp'" alt="" class="w-6 h-6 mt-1 rounded-full" />
                    </a>
                    <a class="text-black font-medium inline-block dark:text-white">
                      {{ Comment.Names }}
                    </a>
                    <div v-if="Comment.Image != ''">
                      <img :src="'/src/assets/images/comment/' + Comment.Image" alt=""
                        class="sm:rounded-lg w-full h-full object-cover" />
                    </div>
                    <div class="sm:px-4 p-2.5 pt-0">
                      <p class="texte">{{ Comment.Content }}</p>
                    </div>
                  </div>
                </div>
              </div>
              <button type="button" class="flex items-center gap-1.5 text-gray-500 hover:text-blue-500 mt-2">
                <ion-icon name="chevron-down-outline"
                  class="ml-auto duration-200 group-aria-expanded:rotate-180"></ion-icon>
                More Comment
              </button>
              <!-- add comment -->
              <div
                class="sm:px-4 sm:py-3 p-2.5 border-t border-gray-100 flex items-center gap-1 dark:border-slate-700/40">
                <img v-if="user.Avatar != ''" :src="'/src/assets/images/avatars/' + user.Avatar" alt=""
                  class="w-9 h-9 rounded-full" />
                <img v-else-if="user.Avatar == ''" :src="'/src/assets/images/avatars/Avatar.webp'" alt=""
                  class="w-9 h-9 rounded-full" />
                <div
                  class="flex-1 bg-slate-100 hover:bg-opacity-80 transition-all rounded-lg cursor-pointer dark:bg-dark3"
                  uk-toggle="target: #create-comment">
                  <div @click="this.RecupId(Post.ID)" class="py-2.5 text-center dark:text-white">
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

    <div class="lg:w-[400px]">
      <div class="lg:space-y-4 lg:pb-8 max-lg:grid sm:grid-cols-2 max-lg:gap-6"
        uk-sticky="media: 1024; end: #js-oversized; offset: 80">
        <!-- group info -->
        <div class="box p-5 px-6">
          <div class="flex items-ce justify-between text-black dark:text-white">
            <h3 class="font-bold text-lg">Info</h3>
          </div>

          <ul class="text-gray-700 space-y-4 mt-2 mb-1 text-sm dark:text-white">
            <li class="flex items-center gap-3">
              <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor"
                class="w-6 h-6">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"
                  d="M12 2a5 5 0 110 10 5 5 0 010-10zm0 0v8m0 0v4m0-4h4m-4 0H8m10 0a5 5 0 00-4 8 5 5 0 004-8zm-6 10v2" />
              </svg>
              <div>
                Description
                <span class="font-semibold text-black dark:text-white">{{
                  sharedData.MygroupProfile.GroupDescription
                  }}</span>
              </div>
            </li>
            <li class="flex items-center gap-3">
              <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
                stroke="currentColor" class="w-6 h-6">
                <path stroke-linecap="round" stroke-linejoin="round"
                  d="M15 19.128a9.38 9.38 0 002.625.372 9.337 9.337 0 004.121-.952 4.125 4.125 0 00-7.533-2.493M15 19.128v-.003c0-1.113-.285-2.16-.786-3.07M15 19.128v.106A12.318 12.318 0 018.624 21c-2.331 0-4.512-.645-6.374-1.766l-.001-.109a6.375 6.375 0 0111.964-3.07M12 6.375a3.375 3.375 0 11-6.75 0 3.375 3.375 0 016.75 0zm8.25 2.25a2.625 2.625 0 11-5.25 0 2.625 2.625 0 015.25 0z" />
              </svg>
              <div>
                Members
                <span v-if="sharedData.MygroupProfile.IdMember" class="font-semibold text-black dark:text-white">
                  {{ sharedData.MygroupProfile.IdMember.length + 1 }} People
                </span>
                <span v-else class="font-semibold text-black dark:text-white">
                  1 People
                </span>
              </div>
            </li>
          </ul>
        </div>

        <div class="box p-5 px-6">
          <div class="flex items-baseline justify-between text-black dark:text-white">
            <h3 class="font-bold text-base">Members</h3>
          </div>
          <div class="side-list">
            <div v-for="user in sharedData.AllUtilisateur" :key="user.ID">
              <div class="side-list-item" v-if="
                (sharedData.MygroupProfile.IdMember &&
                  sharedData.MygroupProfile.IdMember.includes(user.ID)) ||
                sharedData.MygroupProfile.IdCreator == user.ID
              ">
                <router-link v-if="user.ID" :to="{
                  name: 'TimelineProfile',
                  params: { userID: user.ID },
                }"><a>
                    <img v-if="user.Avatar != ''" :src="'/src/assets/images/avatars/' + user.Avatar" alt=""
                      class="side-list-image rounded-full" />
                    <img v-else :src="'/src/assets/images/avatars/Avatar.webp'" alt=""
                      class="side-list-image rounded-full" />
                  </a>
                </router-link>
                <div class="flex-1">
                  <a>
                    <h4 class="side-list-title">
                      {{ user.Nickname }}
                    </h4>
                    <p v-if="sharedData.MygroupProfile.IdCreator == user.ID">
                      Creator
                    </p>
                  </a>
                </div>
              </div>
            </div>
          </div>
        </div>

        <div class="box p-5 px-6">
          <div class="flex items-baseline justify-between text-black dark:text-white">
            <h3 class="font-bold text-base">People you may know</h3>
          </div>
          <div class="side-list">
            <div v-for="user in sharedData.AllUtilisateur" :key="user.ID">
              <div class="side-list-item" v-if="
                (!sharedData.MygroupProfile.IdMember ||
                  !sharedData.MygroupProfile.IdMember.includes(user.ID)) &&
                sharedData.MygroupProfile.IdCreator != user.ID
              ">
                <router-link v-if="user.ID" :to="{
                  name: 'TimelineProfile',
                  params: { userID: user.ID },
                }"><a>
                    <img v-if="user.Avatar != ''" :src="'/src/assets/images/avatars/' + user.Avatar" alt=""
                      class="side-list-image rounded-full" />
                    <img v-else :src="'/src/assets/images/avatars/Avatar.webp'" alt=""
                      class="side-list-image rounded-full" />
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
                <button @click="
                  this.InvitationGroup(
                    sharedData.MygroupProfile.ID_Group,
                    user.ID,
                    $event
                  )
                  " class="button bg-primary-soft text-primary dark:text-white">
                  add
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <Comment></Comment>
</template>

<script lang="js" setup>import sharedData from '../../assets/js/data.js'; import Comment from "../Comment.vue";
</script>


<script lang="js">
import commonMixin from '../../assets/js/untils.js';
export default {
  name: 'Friend',
  mixins: [commonMixin],
};
</script>
<style scoped>
/* Si vous utilisez CSS */
.bg-custom-gray {
  background-color: #d4cdcd;
  font-size: 50%;
  width: max-content;
}
.texte{overflow-wrap: break-word;}
</style>
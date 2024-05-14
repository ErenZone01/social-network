<template>
  <Headers />
  <div id="wrapper">
    <!-- main contents -->
    <main
      id="site__main"
      class="2xl:ml-[--w-side] xl:ml-[--w-side-sm] py-10 p-2.5 h-[calc(100vh-var(--m-top))] mt-[--m-top]"
    >
      <div class="2xl:max-w-[1220px] max-w-[1065px] mx-auto">
        <div class="page-heading">
          <h1 class="page-title">Groups</h1>

          <nav class="nav__underline">
            <ul
              class="group"
              uk-switcher="connect: #group-tabs ; animation: uk-animation-slide-right-medium, uk-animation-slide-left-medium"
            >
              <li><a href="#"> Suggestions </a></li>
              <li><a href="#"> My groups </a></li>
            </ul>
          </nav>
        </div>
        <!-- group list tabs -->
        <div class="uk-switcher" id="group-tabs">
          <!-- card layout 1 -->
          <div class="grid md:grid-cols-4 sm:grid-cols-3 grid-cols-2 gap-2.5">
            <div
              class="card"
              v-for="group in sharedData.UncknowGroup"
              :key="group.ID_Group"
            >
              <a>
                <div class="card-media h-24">
                  <img
                    :src="'/src/assets/images/group/' + group.GroupImage"
                    alt=""
                  />
                  <div class="card-overly"></div>
                </div>
              </a>
              <div class="card-body relative z-10">
                <img
                  v-if="sharedData.AllUtilisateur[group.IdCreator - 1].Avatar"
                  :src="
                    '/src/assets/images/avatars/' +
                    sharedData.AllUtilisateur[group.IdCreator - 1].Avatar
                  "
                  alt=""
                  class="w-10 rounded-full mb-2 shadow -mt-8 relative border-2 border-white dark:border-slate-800"
                />
                <img
                  v-else
                  :src="'/src/assets/images/avatars/Avatar.webp'"
                  alt=""
                  class="w-10 rounded-full mb-2 shadow -mt-8 relative border-2 border-white dark:border-slate-800"
                />
                <a>
                  <h4 class="card-title">{{ group.GroupName }}</h4>
                </a>
                <div class="card-list-info font-normal mt-1">
                  <a href="#"> Travel </a>
                  <div class="md:block hidden">·</div>
                  <div v-if="group.IdMember != null">
                    {{ group.IdMember.length + 1 }} members
                  </div>
                  <div v-else>1 member</div>
                </div>
                <div class="flex gap-2">
                  <button v-if="group.States != 'pending'"
                    type="button"
                    @click="this.InvitationGroup(group.ID_Group, 0)"
                    class="button bg-primary text-white flex-1"
                  >
                    Join
                  </button>
                  <button
                    v-else
                    class="button bg-custom-gray text-gray-700 cursor-not-allowed"
                    disabled
                  >
                    waiting for a response
                  </button>
                  
                </div>
              </div>
              <div :id='"group"+group.ID_Group'></div>
            </div>
          </div>

          <!-- card layout 3 -->
          <div class="grid md:grid-cols-4 sm:grid-cols-3 grid-cols-2 gap-2.5">
            <div
              class="card"
              v-for="group in sharedData.MyGroup"
              :key="group.ID_Group"
            >
              <a>
                <div class="card-media h-24">
                  <img
                    :src="'/src/assets/images/group/' + group.GroupImage"
                    alt=""
                  />
                  <div class="card-overly"></div>
                </div>
              </a>
              <div class="card-body relative z-10">
                <img
                  v-if="sharedData.AllUtilisateur[group.IdCreator - 1].Avatar"
                  :src="
                    '/src/assets/images/avatars/' +
                    sharedData.AllUtilisateur[group.IdCreator - 1].Avatar
                  "
                  alt=""
                  class="w-10 rounded-full mb-2 shadow -mt-8 relative border-2 border-white dark:border-slate-800"
                />
                <img
                  v-else
                  :src="'/src/assets/images/avatars/Avatar.webp'"
                  alt=""
                  class="w-10 rounded-full mb-2 shadow -mt-8 relative border-2 border-white dark:border-slate-800"
                />
                <a>
                  <h4 class="card-title">{{ group.GroupName }}</h4>
                </a>
                <div class="card-list-info font-normal mt-1">
                  <a href="#"> Travel </a>
                  <div class="md:block hidden">·</div>
                  <div v-if="group.IdMember != null">
                    {{ group.IdMember.length + 1 }} members
                  </div>
                  <div v-else>1 member</div>
                </div>
                <div class="flex gap-2">
                  <router-link
                    :to="{
                      name: 'TimelineDiscussion',
                      params: { groupID: group.ID_Group },
                    }"
                    ><a class="button bg-secondery !w-auto"
                      >View</a
                    ></router-link
                  >
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </main>
  </div>
  <Posts></Posts>
</template>

<script lang="js" setup>
import Headers from './Header.vue'
import Posts from './Post.vue'
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
</script>

<style scoped>/* Si vous utilisez CSS */
.bg-custom-gray {
  background-color: #d4cdcd;
  font-size: 100%;
  width: max-content;
}</style>
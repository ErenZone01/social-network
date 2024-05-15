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
            :src='"/src/assets/images/group/group.jpg"'
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
                </div>
              </div>
            </div>
          </div>
        </div>

        

        <div
          class="flex items-center justify-between border-t border-gray-100 px-2 dark:border-slate-700"
        >
          <nav
              class="flex gap-0.5 rounded-xl -mb-px text-gray-600 font-medium text-[15px] dark:text-white max-md:w-full max-md:overflow-x-auto"
            >
              <!-- Lien pour la timeline -->
              <router-link
                :to="{
                  name: 'TimelineDiscussion',
                  params: { groupID: $route.params.groupID },
                }"
                class="inline-block py-3 leading-8 px-3.5 border-b-2 text-blue-600"
                :class="{
                  'border-blue-600': $route.name === 'TimelineDiscussion',
                }"
                >Discussion</router-link
              >

              <!-- Lien pour les amis -->
              <router-link
                :to="{
                  name: 'TimelineChat',
                  params: { groupID: $route.params.groupID },
                }"
                class="inline-block py-3 leading-8 px-3.5 border-b-2 text-blue-600"
                :class="{
                  'border-blue-600': $route.name === 'TimelineChat',
                }"
                >Chat</router-link
              >

              <!-- Lien pour les groupes -->
              <router-link
                :to="{
                  name: 'TimelineEvent',
                  params: { groupID: $route.params.groupID },
                }"
                class="inline-block py-3 leading-8 px-3.5 border-b-2 text-blue-600"
                :class="{ 'border-blue-600': $route.name === 'TimelineEvent' }"
                >Event</router-link
              >
            </nav>
        </div>
      </div>

      <router-view/>
    </div>
  </main>
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
  name: 'TimelineGroup',
  mixins: [commonMixin],
 
  mounted() {
    this.FetchCustomRefGroup()
  }
};
</script>
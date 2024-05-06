<template>
  <!-- create status -->
  <div class="hidden lg:p-20 uk- open" id="create-status" uk-modal="">
    <div
      class="uk-modal-dialog tt relative overflow-hidden mx-auto bg-white shadow-xl rounded-lg md:w-[521px] w-full dark:bg-dark2"
    >
      <div class="text-center py-4 border-b mb-0 dark:border-slate-700">
        <h2 class="text-sm font-medium text-black">Create Post</h2>

        <!-- close button -->
        <button
          type="button"
          class="button-icon absolute top-0 right-0 m-2.5 uk-modal-close"
        >
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
              d="M6 18L18 6M6 6l12 12"
            />
          </svg>
        </button>
      </div>

      <div class="space-y-5 mt-3 p-2">
        <textarea
          class="w-full !text-black placeholder:!text-black !bg-white !border-transparent focus:!border-transparent focus:!ring-transparent !font-normal !text-xl dark:!text-white dark:placeholder:!text-white dark:!bg-slate-800"
          name=""
          id="post"
          rows="6"
          placeholder="What do you have in mind?"
        ></textarea>
      </div>

      <div class="p-5 flex justify-between items-center">
        <div>
          <button
            class="inline-flex items-center py-1 px-2.5 gap-1 font-medium text-sm rounded-full bg-slate-50 border-2 border-slate-100 group aria-expanded:bg-slate-100 aria-expanded: dark:text-white dark:bg-slate-700 dark:border-slate-600"
            type="button"
          >
            <select id="option" v-model="option" name="option">
              <!-- Option "Public" avec selected pour le définir comme par défaut -->
              <option value="Public" selected>Public</option>
              <option value="Allmost private">Allmost private</option>
              <option value="Private">Private</option>
            </select>
            <ion-icon
              name="chevron-down-outline"
              class="text-base duration-500 group-aria-expanded:rotate-180"
            ></ion-icon>
          </button>
        </div>
        <div v-if="option === 'Allmost private'">
          <ul>
            <div v-for="friend in sharedData.MyFollowers" :key="friend.ID">
              <li>
                <input
                  type="checkbox"
                  v-model="selectedFriends"
                  :value="friend.ID"
                />
                <label>{{ friend.Firstname }}</label>
              </li>
            </div>
          </ul>
        </div>
        <div class="flex items-center gap-2">
          <div class="col-span-2">
            <div class="mt-2.5">
              <input
                id="text"
                name="Avatar"
                type="file"
                accept=".jpeg, .jpg, .gif"
                class="!w-full !rounded-lg !bg-transparent !shadow-sm !border-slate-200 dark:!border-slate-800 dark:!bg-white/5"
              />
            </div>
          </div>
        </div>
        <div class="flex items-center gap-2">
          <button
            type="button"
            @click="CreatePost($event)"
            class="button bg-blue-500 text-white py-2 px-12 text-[14px]"
          >
            Create
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="js" setup>
import sharedData from '../assets/js/data.js';
</script>

<script lang="js" >
  import commonMixin from '../assets/js/untils.js';
  export default{
    name: 'Post',
    mixins: [commonMixin],
   data(){ return {option:sharedData.option, selectedFriends:sharedData.selectedFriends}},  
    };

</script>
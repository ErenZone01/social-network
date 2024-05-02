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
              <option value="Public" selected >Public</option>
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
          <div v-for="friend in sharedData.MyFollowings" :key="friend.ID">
            <input
              type="checkbox"
              v-model="selectedFriends"
              :value="friend.ID"
            />
            <label>{{ friend.Firstname }}</label>
          </div>
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
 data(){ return {option:sharedData.option}},
  methods: {
   async CreatePost(e) {
      e.preventDefault();
      let content = document.getElementById("post").value;
      let option = document.getElementsByName("option")[0].value;
      let fileInput = document.getElementsByName("Avatar")[0];
      let image=""//Initialiser le nom de l image a une chaine vide par defaut
      let images = null; // Initialiser les données de l'images à null par défaut
      let byteArrayList = null;
      if (fileInput.files.length > 0) {
        // Vérifier si un fichier a été choisi
        image = fileInput.files[0].name; // Nom du fichier
        images = await fileInput.files[0].arrayBuffer(); // Données de l'images
        // Convertir les données de l'images en tableau de bytes
        let byteArray = new Uint8Array(images);
        byteArrayList = Array.from(byteArray);
      }
      console.log(byteArrayList);
      const post = { Content: content, Privacy: option, ImageData: byteArrayList, Image: image };
     
      fetch("http://localhost:8080/Post", {
        method: "POST",
        body: JSON.stringify(post),
        //ndique que les cookies devraient être inclus dans la requête. Cela est souvent nécessaire lorsqu'une application utilise un système d'authentification basé sur les cookies.
        credentials: "include",
        header: { "Content-Type": "application/json" },
      })
        .then((response) => response.json())
        .then((data) => {
          if (data.Types == "Success"){
            console.log("My post : ", data.Data);
            sharedData.Allpost= data.Data
          }else{

          }
        })
        .catch((error) => console.log("err : ", error));
    },
  },
  };

</script>
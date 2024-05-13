<template>
    <div class="hidden lg:p-20 uk- open" id="create-comment" uk-modal="">
    <div
      class="uk-modal-dialog tt relative overflow-hidden mx-auto bg-white shadow-xl rounded-lg md:w-[521px] w-full dark:bg-dark2"
    >
      <div class="text-center py-4 border-b mb-0 dark:border-slate-700">
        <h2 class="text-sm font-medium text-black">Add Comment</h2>

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
          id="comment"
          rows="6"
          placeholder="What do you have in mind?"
        ></textarea>
      </div>

      <div class="p-5 flex justify-between items-center">
        <div class="flex items-center gap-2">
          <div class="col-span-2">
            <div class="mt-2.5">
              <input
                id="text"
                name="Image"
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
            @click="AddComment($event)"
            class="button bg-blue-500 text-white py-2 px-12 text-[14px]"
          >
            Create
          </button>
        </div>
      </div>
    </div>
  </div>
  </template>
  
  <script lang="js">
  import sharedData from '../assets/js/data.js';
  
  export default {
    name: 'Comment',
    methods: {
     async AddComment(e) {
        e.preventDefault();
        let content = document.getElementById("comment").value;
        let fileInput = document.getElementsByName("Image")[0];
        let image = ""; // Initialize image name to an empty string by default
        let imageData = null; // Initialize image data to null by default
        let byteArrayList = null;
        if (fileInput.files.length > 0) {
          // Check if a file has been selected
          image = fileInput.files[0].name; // Get the file name
          imageData = await fileInput.files[0].arrayBuffer(); // Données de l'images
                  // Convertir les données de l'images en tableau de bytes
        let byteArray = new Uint8Array(imageData);
        byteArrayList = Array.from(byteArray);

        }
      const comment = { Content: content,  ImageData: byteArrayList, Image: image, ID_Post: sharedData.Id};
      fetch("http://localhost:8080/Comment", {
        method: "POST",
        body: JSON.stringify(comment),
        credentials: "include",
        header: { "Content-Type": "application/json" },
      })
      .then((response) => response.json())
        .then((data) => {
          if (data.Types == "Success"){
            console.log("My comment : ", data.Data);
            sharedData.Allcomment= data.Data
          }else{

          }
        })
        .catch((error) => console.log("err : ", error));
      },
    },
  };
  </script>
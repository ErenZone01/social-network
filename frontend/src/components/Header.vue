<template>
  <!-- header -->
  <header
    class="z-[100] h-[--m-top] fixed top-0 left-0 w-full flex items-center bg-white/80 sky-50 backdrop-blur-xl border-b border-slate-200 dark:bg-dark2 dark:border-slate-800">
    <div class="flex items-center w-full xl:px-6 px-2 max-lg:gap-10">
      <div class="2xl:w-[--w-side] lg:w-[--w-side-sm]">
        <!-- left -->
        <div class="flex items-center gap-1">
          <!-- icon menu -->
          <button uk-toggle="target: #site__sidebar ; cls :!-translate-x-0"
            class="flex items-center justify-center w-8 h-8 text-xl rounded-full hover:bg-gray-100 xl:hidden dark:hover:bg-slate-600 group">
            <ion-icon name="menu-outline" class="text-2xl group-aria-expanded:hidden"></ion-icon>
            <ion-icon name="close-outline" class="hidden text-2xl group-aria-expanded:block"></ion-icon>
          </button>
          <div id="logo">
            <router-link to="/Home">
              <a>
                <img src="/src/assets/images/logo.png" alt="" class="w-28 md:block hidden dark:!hidden" />
                <img src="/src/assets/images/logo-light.png" alt="" class="dark:md:block hidden" />
                <img src="/src/assets/images/logo-mobile.png" class="hidden max-md:block w-20 dark:!hidden" alt="" />
                <img src="/src/assets/images/logo-mobile-light.png" class="hidden dark:max-md:block w-20" alt="" />
              </a>
            </router-link>
          </div>
        </div>
      </div>
      <div class="flex-1 relative">
        <div class="max-w-[1220px] mx-auto flex items-center">


          <!-- header icons -->
          <div class="flex items-center sm:gap-4 gap-2 absolute right-5 top-1/2 -translate-y-1/2 text-black">
            <!-- create -->
            <button type="button" class="sm:p-2 p-1 rounded-full relative sm:bg-secondery dark:text-white">
              <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
                stroke="currentColor" class="w-5 h-5 max-sm:hidden">
                <path stroke-linecap="round" stroke-linejoin="round" d="M12 4.5v15m7.5-7.5h-15"></path>
              </svg>
              <ion-icon name="add-circle-outline" class="sm:hidden text-2xl"></ion-icon>
            </button>
            <div
              class="hidden bg-white p-4 rounded-lg overflow-hidden drop-shadow-xl dark:bg-slate-700 md:w-[324px] w-screen border2"
              uk-drop="offset:6;pos: bottom-right; mode: click; animate-out: true; animation: uk-animation-scale-up uk-transform-origin-top-right ">
              <h3 class="font-bold text-md">Create</h3>

              <!-- list -->
              <ul class="uk-slider-items grid-small"
                uk-scrollspy="target: > li; cls: uk-animation-scale-up , uk-animation-slide-right-small; delay: 20 ;repeat: true">
                <li class="w-28 create-btn" @click="toggleForm">
                  <div class="p-3 px-4 rounded-lg bg-purple-100/60 text-purple-600 dark:text-white dark:bg-dark4">
                    <ion-icon name="videocam" class="text-2xl drop-shadow-md"></ion-icon>
                    <div class="mt-1.5 text-sm font-medium">Group</div>
                  </div>
                </li>
              </ul>
            </div>

            <!-- notification -->
            <button type="button" class="sm:p-2 p-1 rounded-full relative sm:bg-secondery dark:text-white"
              uk-tooltip="title: Notification; pos: bottom; offset:6">
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor"
                class="w-6 h-6 max-sm:hidden">
                <path
                  d="M5.85 3.5a.75.75 0 00-1.117-1 9.719 9.719 0 00-2.348 4.876.75.75 0 001.479.248A8.219 8.219 0 015.85 3.5zM19.267 2.5a.75.75 0 10-1.118 1 8.22 8.22 0 011.987 4.124.75.75 0 001.48-.248A9.72 9.72 0 0019.266 2.5z" />
                <path fill-rule="evenodd"
                  d="M12 2.25A6.75 6.75 0 005.25 9v.75a8.217 8.217 0 01-2.119 5.52.75.75 0 00.298 1.206c1.544.57 3.16.99 4.831 1.243a3.75 3.75 0 107.48 0 24.583 24.583 0 004.83-1.244.75.75 0 00.298-1.205 8.217 8.217 0 01-2.118-5.52V9A6.75 6.75 0 0012 2.25zM9.75 18c0-.034 0-.067.002-.1a25.05 25.05 0 004.496 0l.002.1a2.25 2.25 0 11-4.5 0z"
                  clip-rule="evenodd" />
              </svg>
              <div v-if="sharedData.Allnotif"
                class="absolute top-0 right-0 -m-1 bg-red-600 text-white text-xs px-1 rounded-full">
                {{ sharedData.Allnotif.length }}
              </div>
              <ion-icon name="notifications-outline" class="sm:hidden text-2xl"></ion-icon>
            </button>
            <div
              class="hidden bg-white pr-1.5 rounded-lg drop-shadow-xl dark:bg-slate-700 md:w-[365px] w-screen border2"
              uk-drop="offset:6;pos: bottom-right; mode: click; animate-out: true; animation: uk-animation-scale-up uk-transform-origin-top-right ">
              <!-- heading -->
              <div class="flex items-center justify-between gap-2 p-4 pb-2">
                <h3 class="font-bold text-xl">Notifications</h3>

                <div class="flex gap-2.5">
                  <button @click="DeleteAll()" type="button"
                    class="p-1 flex rounded-full focus:bg-secondery dark:text-white">
                    <ion-icon class="text-xl" name="ellipsis-horizontal">Clear All</ion-icon>
                  </button>
                </div>
              </div>

              <div class="text-sm h-[400px] w-full overflow-y-auto pr-2">
                <!-- contents list -->
                <div class="pl-2 p-1 text-sm font-normal dark:text-white">
                  <div v-for="notif in sharedData.Allnotif" :key="notif.ID_Notif">
                    <a href="#"
                      class="relative flex items-center gap-3 p-2 duration-200 rounded-xl pr-10 hover:bg-secondery dark:hover:bg-white/10">
                      <div class="relative w-12 h-12 shrink-0">
                        <img v-if="notif.AvatarSender && notif.AvatarSender != ''" :src="'/src/assets/images/avatars/' + notif.AvatarSender
                          " alt="" class="object-cover w-full h-full rounded-full" />
                        <img v-else :src="'/src/assets/images/avatars/Avatar.webp'" alt=""
                          class="object-cover w-full h-full rounded-full" />
                      </div>
                      <div class="flex-1">
                        <p>
                          <b class="font-bold mr-1"> {{ notif.Sender }} </b>
                          {{ notif.Messages }}
                        </p>
                        <div class="text-xs text-gray-500 mt-1.5 dark:text-white/80">
                          2 hours ago
                        </div>
                      </div>
                      <button v-if="notif.Types != 'event' && notif.States == 'pending'"
                        @click="this.Invitation('true', notif.Types, notif.ID_Notif, $event)" type="button"
                        class="button text-white bg-primary">
                        Accept
                      </button>
                      <button v-if="notif.Types != 'event' && notif.States == 'pending'"
                        @click="this.Invitation('Decline', notif.Types, notif.ID_Notif, $event)" type="button"
                        class="button text-white bg-primary">
                        Decline
                      </button>
                      <button v-if="notif.Types == 'event' && notif.States == 'pending'"
                        @click="this.Invitation('true', notif.Types, notif.ID_Notif, $event)" type="button"
                        class="button text-white bg-primary">
                        Going
                      </button>
                      <button v-if="notif.Types == 'event' && notif.States == 'pending'"
                        @click="this.Invitation('Decline', notif.Types, notif.ID_Notif, $event)" type="button"
                        class="button text-white bg-primary">
                        Not Going
                      </button>
                    </a>
                  </div>
                </div>
              </div>

              <!-- footer -->
              <a href="#">
                <div
                  class="text-center py-4 border-t border-slate-100 text-sm font-medium text-blue-600 dark:text-white dark:border-gray-600">
                  View Notifications
                </div>
              </a>

              <div
                class="w-3 h-3 absolute -top-1.5 right-3 bg-white border-l border-t rotate-45 max-md:hidden dark:bg-dark3 dark:border-transparent">
              </div>
            </div>

            <!-- profile -->
            <div class="rounded-full relative bg-secondery cursor-pointer shrink-0">
              <img v-if="
                sharedData.Myaccount.Avatar &&
                sharedData.Myaccount.Avatar != ''
              " :src="'/src/assets/images/avatars/' + sharedData.Myaccount.Avatar
                  " alt="" class="sm:w-9 sm:h-9 w-7 h-7 rounded-full shadow shrink-0" />
              <img v-else :src="'/src/assets/images/avatars/Avatar.webp'" alt=""
                class="sm:w-9 sm:h-9 w-7 h-7 rounded-full shadow shrink-0" />
            </div>
            <div class="hidden bg-white rounded-lg drop-shadow-xl dark:bg-slate-700 w-64 border2"
              uk-drop="offset:6;pos: bottom-right;animate-out: true; animation: uk-animation-scale-up uk-transform-origin-top-right ">
              <a>
                <router-link v-if="sharedData.Myaccount.ID" :to="{
                  name: 'TimelineProfile',
                  params: { userID: sharedData.Myaccount.ID },
                }">
                  <div class="p-4 py-5 flex items-center gap-4">
                    <img v-if="
                      sharedData.Myaccount.Avatar &&
                      sharedData.Myaccount.Avatar != ''
                    " :src="'/src/assets/images/avatars/' +
                        sharedData.Myaccount.Avatar
                        " alt="" class="w-10 h-10 rounded-full shadow" />
                    <img v-else :src="'/src/assets/images/avatars/Avatar.webp'" alt=""
                      class="w-10 h-10 rounded-full shadow" />
                    <div class="flex-1">
                      <h4 class="text-sm font-medium text-black">
                        {{ sharedData.Myaccount.Nickname }}
                      </h4>
                      <div class="text-sm mt-1 text-blue-600 font-light dark:text-white/70">
                        {{ sharedData.Myaccount.Firstname }}
                        {{ sharedData.Myaccount.Lastname }}
                      </div>
                    </div>
                  </div>
                </router-link>
              </a>

              <hr class="dark:border-gray-600/60" />

              <nav class="p-2 text-sm text-black font-normal dark:text-white">
                <a @click="Decon($event)">
                  <div
                    class="flex items-center gap-2.5 hover:bg-secondery p-2 px-2.5 rounded-md dark:hover:bg-white/10">
                    <svg class="w-6" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"
                      stroke="currentColor">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                        d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1">
                      </path>
                    </svg>
                    Log Out
                  </div>
                </a>
              </nav>
            </div>

            <div class="flex items-center gap-2 hidden">
              <img src="/src/assets/images/avatars/avatar-2.jpg" alt="" class="w-9 h-9 rounded-full shadow" />

              <div class="w-20 font-semibold text-gray-600">Hamse</div>

              <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
                stroke="currentColor" class="w-5 h-5">
                <path stroke-linecap="round" stroke-linejoin="round" d="M19.5 8.25l-7.5 7.5-7.5-7.5" />
              </svg>
            </div>
          </div>
        </div>
      </div>
    </div>
  </header>

  <!-- sidebar -->
  <div id="site__sidebar"
    class="fixed top-0 left-0 z-[99] pt-[--m-top] overflow-hidden transition-transform xl:duration-500 max-xl:w-full max-xl:-translate-x-full">
    <!-- sidebar inner -->
    <div
      class="p-2 max-xl:bg-white shadow-sm 2xl:w-72 sm:w-64 w-[80%] h-[calc(100vh-64px)] relative z-30 max-lg:border-r dark:max-xl:!bg-slate-700 dark:border-slate-700">
      <div class="pr-4">
        <nav id="side">
          <ul>
            <li :class="{
              active: $route.name === 'Home',
            }">
              <router-link to="/Home">
                <a>
                  <img src="/src/assets/images/icons/home.png" alt="feeds" class="w-6" />
                  <span> Feed </span>
                </a>
              </router-link>
            </li>
            <li :class="{
              active: $route.name === 'Messages',
            }">
              <router-link to="/Messages">
                <a>
                  <img src="/src/assets/images/icons/message.png" alt="messages" class="w-5" />
                  <span> messages </span>
                </a>
              </router-link>
            </li>
            <li :class="{
              active: $route.name === 'Group',
            }">
              <router-link to="/Group">
                <a>
                  <img src="/src/assets/images/icons/group.png" alt="groups" class="w-6" />
                  <span> Groups </span>
                </a>
              </router-link>
            </li>
          </ul>
        </nav>
      </div>
    </div>
  </div>

  <!-- group creation form  -->
  <div class="group-form-container" v-show="showForm">
    <button class="close-button" @click="closeForm">×</button>
    <form class="group-form">
      <div class="form-group">
        <label for="groupName">Group Name</label>
        <input type="text" id="groupName" v-model="groupName" />
      </div>
      <div class="form-group">
        <label for="groupDescription">Group Description</label>
        <textarea id="groupDescription" v-model="groupDescription"></textarea>
      </div>
      <button @click="this.CreateGroup($event)">Create</button>
    </form>
  </div>
  <div class="overlay" v-if="showForm"></div>
</template>

<script lang="js" setup>
import sharedData from '../assets/js/data.js';</script>

<script lang="js">
import commonMixin from '../assets/js/untils.js';
import { CustomFetch } from '../assets/js/untils.js';

import { ref } from "vue";
// group creation
const showForm = ref(false);
const groupName = ref("");
const groupDescription = ref("");

export default {
  name: 'Header',
  mixins: [commonMixin],
  async mounted() {
    this.GetData();
  },
  methods: {
    async CreateGroup(event) {
      event.preventDefault();
      let avatar = ""; // Initialiser le nom de l'avatar à une chaîne vide par défaut
      let byteArrayList = null;
      var StructGroup = {
        GroupName: groupName.value,
        GroupImage: avatar,
        GroupImageFile: byteArrayList,
        GroupDescription: groupDescription.value,
        IdCreator: 0
      }

      var fetch = await CustomFetch("http://localhost:8080/Group", "POST", StructGroup);
      if (fetch.Types == "Success") {
        console.log("Group created");
        sharedData.UncknowGroup = fetch.Data.UncknowGroup;
        sharedData.MyGroup = fetch.Data.MyGroup; toggleForm(); closeForm();
      } else {
        console.log("Error of Creation Group : ", fetch.Msg);
        this.$router.push("/Login");
      }
      console.log("Name Group : ", groupName.value, ", Description : ", groupDescription.value, ", Image : ", avatar)

    },
    async DeleteAll(){
      var fetch = await CustomFetch("http://localhost:8080/Delete", "GET", null)
      if (fetch.Types == "Success") {
        sharedData.Allnotif= fetch.Data
      }else{
        console.log("Error of Delete : ", fetch.Msg);
        this.$router.push("/Login");
      } 
    }
  }
};


// Function to toggle the display of the form
const toggleForm = () => {
  showForm.value = !showForm.value;
};
const closeForm = () => {
  showForm.value = false;
  groupName.value = "";
  groupDescription.value = "";
};

// Gérer le changement de la photo de profil
const handleProfilePicChange = (event) => {
  const file = event.target.files[0];
  groupProfilePic.value = file;
  groupProfilePicName = file.Name
};

</script>

<style scoped>
.overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(250, 245, 245, 0.712);
  z-index: 999;
}

.header {
  backdrop-filter: blur(100px);
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100vw;
  height: 50px;
  position: fixed;
  background-color: #476678;
  padding: 15px 0;
}

.logo {
  width: 30%;
}

.logo img {
  width: 100px !important;
}

nav {
  padding: 12px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 70%;
  cursor: pointer;
}

nav img {
  width: 25px;
  height: 25px;
}

.pages {
  width: 55%;
  display: flex;
  justify-content: space-around;
  align-items: center;
  width: 50%;
}

.profil,
.home {
  display: flex;
  justify-content: space-around;
  align-items: center;
  width: 100px;
}

.msg-log {
  display: flex;
  justify-content: space-around;
  align-items: center;
  width: 35%;
}

.messagerie {
  display: flex;
  justify-content: space-around;
  align-items: center;
  width: 70%;
}

.logOut {
  display: flex;
  align-items: center;
  font-family: Arial, Helvetica, sans-serif;
  font-weight: 800;
}

.logOut img {
  margin-left: 5px;
  width: 20px;
}

/* responsivite */
@media screen and (max-width: 800px) {
  nav span {
    display: none;
  }
}

.group-form-container {
  z-index: 1000;
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  background-color: #f9f9f9;
  padding: 30px;
  border-radius: 15px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  max-height: 550px;
  overflow-y: auto;
  max-width: 90%;
}

.group-form {
  max-width: 500px;
  margin: 0 auto;
}

.form-group {
  margin-bottom: 25px;
}

.form-group label {
  font-size: 16px;
  color: #333;
}

.group-form input,
.group-form textarea,
.group-form select {
  margin-top: 5px;
  width: 95%;
  padding: 10px;
  border: 1px solid #ccc;
  /* Bordure */
  border-radius: 5px;
  font-size: 14px;
  transition: border-color 0.3s ease;
  /* Transition de la couleur de la bordure */
}

.group-form select {
  width: 100% !important;
}

.group-form textarea {
  resize: none;
}

.group-form input:focus,
.group-form textarea:focus,
.group-form select:focus {
  border-color: #5c9edd;
  /* Couleur de la bordure au focus */
  outline: none;
  /* Supprimer le contour */
}

.group-form button {
  width: 100%;
  padding: 12px;
  background-color: #5c9edd;
  /* Couleur de fond du bouton */
  color: #fff;
  /* Couleur du texte du bouton */
  border: none;
  border-radius: 5px;
  font-size: 16px;
  cursor: pointer;
  transition: background-color 0.3s ease;
  /* Transition de la couleur de fond */
}

.group-form button:hover {
  background-color: #4b89d0;
}

.close-button {
  position: absolute;
  top: -15px;
  right: 15px;
  background-color: transparent;
  border: none;
  font-size: 40px;
  cursor: pointer;
  width: 30px;
  height: 30px;
  color: gray;
}

.close-button:hover {
  color: #cc0000;
  background-color: transparent;
}

.create-btn {
  cursor: pointer;
}
</style>
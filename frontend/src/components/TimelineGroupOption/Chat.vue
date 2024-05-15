<template>
  <!-- message center -->
  <div class="flex-1">

    <!-- chat heading -->
    <div
      class="flex items-center justify-between gap-2 w- px-6 py-3.5 z-10 border-b dark:border-slate-700 uk-animation-slide-top-medium">

      <div class="flex items-center sm:gap-4 gap-2">

        <!-- toggle for mobile -->
        <button type="button" class="md:hidden" uk-toggle="target: #side-chat ; cls: max-md:-translate-x-full">
          <ion-icon name="chevron-back-outline" class="text-2xl -ml-4"></ion-icon>
        </button>

        <div>
          <!-- <img src="/src/assets/images/avatars/avatar-6.jpg" alt="" class="w-8 h-8 rounded-full shadow"> -->
          <img v-if="
            sharedData.MygroupProfile.GroupImage &&
            sharedData.MygroupProfile.GroupImage != ''
          " :src="'/src/assets/images/group/' +
            sharedData.MygroupProfile.GroupImage
            " class="w-8 h-8 rounded-full shadow" alt="" />
          <img v-else :src="'/src/assets/images/avatars/Avatar.webp'" alt="" class="w-8 h-8 rounded-full shadow" />
          <div class="w-2 h-2 bg-teal-500 rounded-full absolute right-0 bottom-0 m-px"></div>
        </div>
        <div>
          <div class="text-base font-bold">{{ sharedData.MygroupProfile.GroupName }} (Chat Group)</div>
          <div class="text-xs text-green-500 font-semibold"> Active </div>
        </div>
      </div>

    </div>

    <!-- chats bubble -->
    <div class="w-full p-5 py-10 overflow-y-auto md:h-[calc(100vh-204px)] h-[calc(100vh-195px)]">
      <!-- v-if="IfAnAccountFollowMe(sharedData.LatestChat) || IfIFollowAnAccount(sharedData.LatestChat)" -->


      <!-- target -->
      <div v-for="chat in sharedData.AllChatsGroup" :key="chat.ID_User" class="text-sm font-medium space-y-6">
        <!-- received -->
        <div v-if="chat.ID_User != sharedData.Myaccount.ID" class="flex gap-3">
          <div v-for="user in sharedData.AllUsers" :key="user.Id" class="flex flex-col">
            <div v-if="user.ID == chat.ID_User">
              <img v-if="user.Avatar != ''" :src="'/src/assets/images/avatars/' + user.Avatar
                " alt="" class="w-9 h-9 rounded-full shadow" />
              <img v-else-if="user.Avatar == ''" :src="'/src/assets/images/avatars/Avatar.webp'" alt=""
                class="w-9 h-9 rounded-full shadow" />
              <span>{{ user.Nickname }}</span>
            </div>
          </div>
          <div class="px-4 py-2 rounded-[20px] max-w-sm bg-secondery">{{ chat.Content }}</div>
        </div>

        <!-- sent -->
        <div v-else class="flex gap-2 flex-row-reverse items-end">
          <!-- <img src="/src/assets/images/avatars/avatar-3.jpg" alt="" class="w-5 h-5 rounded-full shadow"> -->
          <img v-if="sharedData.Myaccount.Avatar != ''" :src="'/src/assets/images/avatars/' + sharedData.Myaccount.Avatar
            " alt="" class="w-5 h-5 rounded-full shadow" />
          <img v-else-if="sharedData.Myaccount.Avatar == ''" :src="'/src/assets/images/avatars/Avatar.webp'" alt=""
            class="w-5 h-5 rounded-full shadow" />
          <div class="px-4 py-2 rounded-[20px] max-w-sm bg-gradient-to-tr from-sky-500 to-blue-500 text-white shadow">
            {{ chat.Content }} </div>
        </div>
        <br />
      </div>
      <!-- end -->
    </div>

    <!-- sending message area -->
    <div class="flex items-center md:gap-4 gap-2 md:p-3 p-2 overflow-hidden">

      <div id="message__wrap" class="flex items-center gap-2 h-full dark:text-white -mt-1.5">

        <button type="button" class="shrink-0">
          <ion-icon class="text-3xl flex" name="add-circle-outline"></ion-icon>
        </button>
        <div
          class="dropbar pt-36 h-60 bg-gradient-to-t via-white from-white via-30% from-30% dark:from-slate-900 dark:via-900"
          uk-drop="stretch: x; target: #message__wrap ;animation:  slide-bottom ;animate-out: true; pos: top-left; offset:10 ; mode: click ; duration: 200">

          <div class="sm:w-full p-3 flex justify-center gap-5"
            uk-scrollspy="target: > button; cls: uk-animation-slide-bottom-small; delay: 100;repeat:true">

            <button type="button"
              class="bg-sky-50 text-sky-600 border border-sky-100 shadow-sm p-2.5 rounded-full shrink-0 duration-100 hover:scale-[1.15] dark:bg-dark3 dark:border-0">
              <ion-icon class="text-3xl flex" name="image"></ion-icon>
            </button>
            <button type="button"
              class="bg-green-50 text-green-600 border border-green-100 shadow-sm p-2.5 rounded-full shrink-0 duration-100 hover:scale-[1.15] dark:bg-dark3 dark:border-0">
              <ion-icon class="text-3xl flex" name="images"></ion-icon>
            </button>
            <button type="button"
              class="bg-pink-50 text-pink-600 border border-pink-100 shadow-sm p-2.5 rounded-full shrink-0 duration-100 hover:scale-[1.15] dark:bg-dark3 dark:border-0">
              <ion-icon class="text-3xl flex" name="document-text"></ion-icon>
            </button>
            <button type="button"
              class="bg-orange-50 text-orange-600 border border-orange-100 shadow-sm p-2.5 rounded-full shrink-0 duration-100 hover:scale-[1.15] dark:bg-dark3 dark:border-0">
              <ion-icon class="text-3xl flex" name="gift"></ion-icon>
            </button>


          </div>

        </div>

        <button type="button" class="shrink-0">
          <ion-icon class="text-3xl flex" name="happy-outline">😊</ion-icon>
        </button>
        <div class="dropbar p-2"
          uk-drop="stretch: x; target: #message__wrap ;animation: uk-animation-scale-up uk-transform-origin-bottom-left ;animate-out: true; pos: top-left ; offset:2; mode: click ; duration: 200 ">
          <div class="sm:w-60 bg-white shadow-lg border rounded-xl  pr-0 dark:border-slate-700 dark:bg-dark3">

            <h4 class="text-sm font-semibold p-3 pb-0">Send Imogi</h4>

            <div class="grid grid-cols-5 overflow-y-auto max-h-44 p-3 text-center text-xl">

              <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200"
                @click="EmojiClicked('😊')"> 😊 </div>
              <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200"
                @click="EmojiClicked('🤩')"> 🤩 </div>
              <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200"
                @click="EmojiClicked('😎')"> 😎 </div>
              <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200"
                @click="EmojiClicked('😊')"> 😊 </div>
              <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200"
                @click="EmojiClicked('🤩')"> 🤩 </div>
              <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200"
                @click="EmojiClicked('😎')"> 😎 </div>
              <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200"
                @click="EmojiClicked('😊')"> 😊 </div>
              <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200"
                @click="EmojiClicked('🤩')"> 🤩 </div>
              <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200"
                @click="EmojiClicked('😎')"> 😎 </div>
              <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200"
                @click="EmojiClicked('😊')"> 😊 </div>
              <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200"
                @click="EmojiClicked('🤩')"> 🤩 </div>
              <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200"
                @click="EmojiClicked('😎')"> 😎 </div>
              <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200"
                @click="EmojiClicked('😊')"> 😊 </div>
              <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200"
                @click="EmojiClicked('🤩')"> 🤩 </div>
              <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200"
                @click="EmojiClicked('😎')"> 😎 </div>

            </div>


          </div>


        </div>

      </div>

      <div class="relative flex-1">
        <textarea placeholder="Write your message" rows="1"
          class="w-full resize-none bg-secondery rounded-full px-4 p-2" id="mes"
          @keydown.enter="SendAndDisplay($event)"></textarea>
        <!-- <button type="text" >send</button> -->
        <button type="button" class="text-white shrink-0 p-2 absolute right-0.5 top-0">
          <ion-icon class="text-xl flex" name="send-outline"></ion-icon>
        </button>

      </div>

      <button type="button" class="flex h-full dark:text-white">
        <ion-icon class="text-3xl flex -mt-3" name="heart-outline"></ion-icon>
      </button>

    </div>

  </div>
</template>

<script lang="js" setup>
import sharedData from '../../assets/js/data.js';
import { mysocket } from "../../assets/js/socket.js";
import { CustomFetch } from "../../assets/js/untils.js";

</script>


<script lang="js">
import commonMixin from '../../assets/js/untils.js';
export default {
  mixins: [commonMixin],
  async mounted() { this.FetchCustomRefGroup() },
  watch: {
    '$route'(to, from) {
      // Appeler fetchData() lorsque la route change
      this.FetchCustomRefGroup();
    }
  },
  methods: {

    EmojiClicked(emoji) {
      let textarea = document.getElementById("mes");
      textarea.value += emoji
      // console.log('Emoji choisi :', emoji);
    },

    async SendAndDisplay(event) {
      if (event.keyCode === 13) {
        // Empêche le saut de ligne par défaut dans le textarea
        event.preventDefault();
        let textarea = document.getElementById("mes");
        if (textarea.value.trim() === "") {
          textarea.value = ""
          return
        }
        var groupID = parseInt(this.$route.params.groupID);
        // var groupID = this.$route.params.groupID
        // console.log("le params est : ", groupID)
        const chat = { ID_Receiver: groupID, Content: textarea.value, Types: "group", ID_User: sharedData.Myaccount.ID };
        // console.log("chat", chat);
        // sharedData.AllChats=[]
        const data = await CustomFetch("http://localhost:8080/Chat", "POST", chat)
        console.log("DAta :", data.Data)
        sharedData.AllChatsGroup = data.Data;
        // sharedData.AllChats[data.Data.length - 1].ID_User
        const render = { Payload: "group", To: chat.ID_Receiver, Obj: sharedData.AllChatsGroup, From: chat.ID_User }
        mysocket.send(render)
        textarea.value = "";
      }
    }
  }



};
</script>

<style scoped></style>
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
            sharedData.LatestChat.Avatar &&
            sharedData.LatestChat.Avatar != ''
          " :src="'/src/assets/images/avatars/' +
            sharedData.LatestChat.Avatar
            " class="w-8 h-8 rounded-full shadow" alt="" />
          <img v-else :src="'/src/assets/images/avatars/Avatar.webp'" alt="" class="w-8 h-8 rounded-full shadow" />
          <div class="w-2 h-2 bg-teal-500 rounded-full absolute right-0 bottom-0 m-px"></div>
        </div>
        <div>
          <div v-if="IfAnAccountFollowMe(sharedData.LatestChat) || IfIFollowAnAccount(sharedData.LatestChat)"
            class="text-base font-bold">{{ sharedData.LatestChat.Firstname }} {{
              sharedData.LatestChat.Lastname }}</div>
          <div v-else class="text-base font-bold"> JOKKO
          </div>
          <div class="text-xs text-green-500 font-semibold"> Online</div>
        </div>

      </div>

    </div>

    <!-- chats bubble -->
    <div class="w-full p-5 py-10 overflow-y-auto md:h-[calc(100vh-204px)] h-[calc(100vh-195px)]">
      <!-- v-if="IfAnAccountFollowMe(sharedData.LatestChat) || IfIFollowAnAccount(sharedData.LatestChat)" -->
      <div class="py-10 text-center text-sm lg:pt-8"
        v-if="IfAnAccountFollowMe(sharedData.LatestChat) || IfIFollowAnAccount(sharedData.LatestChat)">
        <!-- <img src="/src/assets/images/avatars/avatar-6.jpg" class="w-24 h-24 rounded-full mx-auto mb-3" alt=""> -->
        <img v-if="
          sharedData.LatestChat.Avatar &&
          sharedData.LatestChat.Avatar != ''
        " :src="'/src/assets/images/avatars/' +
          sharedData.LatestChat.Avatar
          " class="w-24 h-24 rounded-full mx-auto mb-3" alt="" />
        <img v-else :src="'/src/assets/images/avatars/Avatar.webp'" alt=""
          class="w-24 h-24 rounded-full mx-auto mb-3" />
        <div class="mt-8">
          <div class="md:text-xl text-base font-medium text-black dark:text-white">{{
            sharedData.LatestChat.Firstname }} {{ sharedData.LatestChat.Lastname }}</div>
          <div class="text-gray-500 text-sm   dark:text-white/80"> @{{ sharedData.LatestChat.Nickname }}</div>
        </div>
        <div class="mt-3.5">
          <a class="inline-block rounded-lg px-4 py-1.5 text-sm font-semibold bg-secondery">
            <router-link :to="{
              name: 'TimelineProfile',
              params: { userID: sharedData.LatestChat.ID },
            }"> View profile</router-link>
          </a>
        </div>
      </div>
      <div v-else class="py-10 text-center text-sm lg:pt-8">
        <img src="/src/assets/images/avatars/avatar-6.jpg" class="w-24 h-24 rounded-full mx-auto mb-3" alt="">
        <div class="mt-8">
          <div class="md:text-xl text-base font-medium text-black dark:text-white">JOKKO </div>
          <div class="text-gray-500 text-sm   dark:text-white/80"> @Jokko</div>
        </div>
        <div class="mt-3.5">
          <a href="timeline.html" class="inline-block rounded-lg px-4 py-1.5 text-sm font-semibold bg-secondery">View
            profile</a>
        </div>
      </div>
      <!-- target -->
      <div v-for="chat in sharedData.AllChats" :key="chat.ID_User" class="text-sm font-medium space-y-6"
        v-if="IfAnAccountFollowMe(sharedData.LatestChat) || IfIFollowAnAccount(sharedData.LatestChat)">

        <!-- received -->
        <div v-if="chat.ID_User != sharedData.Myaccount.ID" class="flex gap-3">
          <!-- <img src="/src/assets/images/avatars/avatar-2.jpg" alt="" class="w-9 h-9 rounded-full shadow"> -->
          <img v-if="sharedData.LatestChat.Avatar != ''" :src="'/src/assets/images/avatars/' + sharedData.LatestChat.Avatar
            " alt="" class="w-9 h-9 rounded-full shadow" />
          <img v-else-if="sharedData.LatestChat.Avatar == ''" :src="'/src/assets/images/avatars/Avatar.webp'" alt=""
            class="w-9 h-9 rounded-full shadow" />
          <div class="px-4 py-2 rounded-[20px] max-w-sm bg-secondery">{{ chat.Content }} </div>
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

        <!-- time -->
        <!-- <div class="flex justify-center ">
  <div class="font-medium text-gray-500 text-sm dark:text-white/70">
    April 8,2023,6:30 AM
  </div>
</div>

<!-- received -->
        <!-- <div class="flex gap-3">
  <img src="/src/assets/images/avatars/avatar-2.jpg" alt="" class="w-9 h-9 rounded-full shadow">
  <div class="px-4 py-2 rounded-[20px] max-w-sm bg-secondery"> I’m selling a photo of a sunset. It’s a
    print on canvas, signed by the photographer. Do you like it? 😊 </div>
</div> -->

        <!-- sent -->
        <!-- <div class="flex gap-2 flex-row-reverse items-end">
  <img src="/src/assets/images/avatars/avatar-3.jpg" alt="" class="w-4 h-4 rounded-full shadow">
  <div
    class="px-4 py-2 rounded-[20px] max-w-sm bg-gradient-to-tr from-sky-500 to-blue-500 text-white shadow">
    Wow, it’s beautiful. How much ? 😍 </div>
</div> -->

        <!-- sent media-->
        <!-- <div class="flex gap-2 flex-row-reverse items-end">
  <img src="/src/assets/images/avatars/avatar-3.jpg" alt="" class="w-4 h-4 rounded-full shadow">

  <a class="block rounded-[18px] border overflow-hidden" href="#">
    <div class="max-w-md">
      <div class="max-w-full relative w-72">
        <div class="relative" style="padding-bottom: 57.4286%">
          <div class="w-full h-full absolute inset-0">
            <img src="/src/assets/images/product/product-1.jpg" alt=""
              class="block max-w-full max-h-52 w-full h-full Groupsobject-cover">
          </div>
        </div>
      </div>
    </div>
  </a>

</div> -->

        <!-- time -->
        <!-- <div class="flex justify-center ">
  <div class="font-medium text-gray-500 text-sm dark:text-white/70">
    April 8,2023,6:30 AM
  </div>
</div> -->


        <!-- received -->
        <!-- <div class="flex gap-3">
  <img src="/src/assets/images/avatars/avatar-2.jpg" alt="" class="w-9 h-9 rounded-full shadow">
  <div class="px-4 py-2 rounded-[20px] max-w-sm bg-secondery"> I’m glad you like it. I’m asking for $200
    🤑</div>
</div> -->

        <!-- sent -->
        <!-- <div class="flex gap-2 flex-row-reverse items-end">
  <img src="/src/assets/images/avatars/avatar-3.jpg" alt="" class="w-5 h-5 rounded-full shadow">
  <div
    class="px-4 py-2 rounded-[20px] max-w-sm bg-gradient-to-tr from-sky-500 to-blue-500 text-white shadow">
    $200? Too steep. Can you lower the price a bit? 😕</div>
</div> -->

        <!-- received -->
        <!-- <div class="flex gap-3">
  <img src="/src/assets/images/avatars/avatar-2.jpg" alt="" class="w-9 h-9 rounded-full shadow">
  <div class="px-4 py-2 rounded-[20px] max-w-sm bg-secondery"> Well, I can’t go too low because I paid a
    lot. But I’m willing to negotiate. What’s your offer? 🤔 </div>

</div> -->

        <!-- sent -->
        <!-- <div class="flex gap-2 flex-row-reverse items-end">
  <img src="/src/assets/images/avatars/avatar-3.jpg" alt="" class="w-5 h-5 rounded-full shadow">
  <div
    class="px-4 py-2 rounded-[20px] max-w-sm bg-gradient-to-tr from-sky-500 to-blue-500 text-white shadow">
    Sorry, can’t pay more than $150. 😅</div>
</div> -->

        <!-- time -->
        <!-- <div class="flex justify-center ">
  <div class="font-medium text-gray-500 text-sm dark:text-white/70">
    April 8,2023,6:30 AM
  </div>
</div> -->

        <!-- received -->
        <!-- <div class="flex gap-3">
  <img src="/src/assets/images/avatars/avatar-2.jpg" alt="" class="w-9 h-9 rounded-full shadow">
  <div class="px-4 py-2 rounded-[20px] max-w-sm bg-secondery"> $150? Too low. Photo worth more. 😬</div>
</div> -->

        <!-- sent -->
        <!-- <div class="flex gap-2 flex-row-reverse items-end">
  <img src="/src/assets/images/avatars/avatar-3.jpg" alt="" class="w-5 h-5 rounded-full shadow">
  <div
    class="px-4 py-2 rounded-[20px] max-w-sm bg-gradient-to-tr from-sky-500 to-blue-500 text-white shadow">
    Too high. I Can’t . How about $160? Final offer. 😬 </div>
</div> -->

        <!-- received -->
        <!-- <div class="flex gap-3">
  <img src="/src/assets/images/avatars/avatar-2.jpg" alt="" class="w-9 h-9 rounded-full shadow">
  <div class="px-4 py-2 rounded-[20px] max-w-sm bg-secondery"> Fine, fine. You’re hard to please. I’ll
    take $160, but only because I like you. 😍</div>
</div> -->

        <!-- sent -->
        <!-- <div class="flex gap-2 flex-row-reverse items-end">
  <img src="/src/assets/images/avatars/avatar-3.jpg" alt="" class="w-5 h-5 rounded-full shadow">
  <div
    class="px-4 py-2 rounded-[20px] max-w-sm bg-gradient-to-tr from-sky-500 to-blue-500 text-white shadow">
    Great, thank you. I appreciate it. I love this photo and can’t wait to hang it. 😩 </div>
</div> -->

      </div>
      <div v-else class="text-sm font-medium space-y-6">
        <!-- received -->
        <div class="flex gap-3">
          <img src="/src/assets/images/avatars/avatar-6.jpg" alt="" class="w-9 h-9 rounded-full shadow">
          <div class="px-4 py-2 rounded-[20px] max-w-sm bg-secondery"> Envie de discutter? 😊 Selectionnez un de
            vos abonnés/abonnement! </div>
        </div>
      </div>
      <!-- end -->
    </div>

    <!-- sending message area -->
    <div class="flex items-center md:gap-4 gap-2 md:p-3 p-2 overflow-hidden"
      v-if="IfAnAccountFollowMe(sharedData.LatestChat) || IfIFollowAnAccount(sharedData.LatestChat)">

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
          <ion-icon class="text-3xl flex" name="happy-outline"></ion-icon>
        </button>
        <div class="dropbar p-2"
          uk-drop="stretch: x; target: #message__wrap ;animation: uk-animation-scale-up uk-transform-origin-bottom-left ;animate-out: true; pos: top-left ; offset:2; mode: click ; duration: 200 ">
          <div class="sm:w-60 bg-white shadow-lg border rounded-xl  pr-0 dark:border-slate-700 dark:bg-dark3">

            <h4 class="text-sm font-semibold p-3 pb-0">Send Imogi</h4>

            <div class="grid grid-cols-5 overflow-y-auto max-h-44 p-3 text-center text-xl">

              <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200"> 😊
              </div>
              <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200"> 🤩
              </div>
              <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200"> 😎
              </div>
              <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200"> 🥳
              </div>
              <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200"> 😂
              </div>
              <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200"> 🥰
              </div>
              <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200"> 😡
              </div>
              <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200"> 😊
              </div>
              <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200"> 🤩
              </div>
              <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200"> 😎
              </div>
              <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200"> 🥳
              </div>
              <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200"> 😂
              </div>
              <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200"> 🥰
              </div>
              <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200"> 😡
              </div>
              <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200"> 🤔
              </div>
              <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200"> 😊
              </div>
              <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200"> 🤩
              </div>
              <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200"> 😎
              </div>
              <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200"> 🥳
              </div>
              <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200"> 😂
              </div>

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
import { CustomFetch } from "../../assets/js/untils.js";

</script>


<script lang="js">
import commonMixin from '../../assets/js/untils.js';
export default {
  mixins: [commonMixin],


  async SendAndDisplay(event) {
    if (event.keyCode === 13) {
      // Empêche le saut de ligne par défaut dans le textarea
      event.preventDefault();
      let textarea = document.getElementById("mes");
      if (textarea.value.trim() === "") {
        textarea.value = ""
        return
      }
      const chat = { ID_Receiver: sharedData.LatestChat.ID, Content: textarea.value, Types: "chat", ID_User: sharedData.Myaccount.ID };
      console.log("chat", chat);
      // sharedData.AllChats=[]
      const data = await CustomFetch("http://localhost:8080/Chat", "POST", chat)
      sharedData.AllChats = data.Data;
      // sharedData.AllChats[data.Data.length - 1].ID_User
      const render = { Payload: "chat", To: chat.ID_Receiver, Obj: sharedData.AllChats, From: chat.ID_User }
      mysocket.send(render)
      textarea.value = "";
      // Traitez votre action ici, par exemple, envoyez le message
    }
  }
};
</script>

<style scoped></style>
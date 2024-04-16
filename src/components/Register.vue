<template>
  <div class="sm:flex">
    <div
      class="relative lg:w-[580px] md:w-96 w-full p-10 min-h-screen bg-white shadow-xl flex items-center pt-10 dark:bg-slate-900 z-10"
    >
      <div
        class="w-full lg:max-w-sm mx-auto space-y-10"
        uk-scrollspy="target: > *; cls: uk-animation-scale-up; delay: 100 ;repeat: true"
      >
        <!-- logo image-->
        <a href="#">
          <img
            src="/src/assets/images/logo.png"
            class="w-28 absolute top-10 left-10 dark:hidden"
            alt=""
        /></a>
        <a href="#">
          <img
            src="/src/assets/images/logo-light.png"
            class="w-28 absolute top-10 left-10 hidden dark:!block"
            alt=""
        /></a>
        <!-- logo icon optional -->
        <div class="hidden">
          <img
            class="w-12"
            src="/src/assets/images/logo-icon.png"
            alt="Socialite html template"
          />
        </div>
        <!-- title -->
        <div>
          <h2 class="text-2xl font-semibold mb-1.5">Sign up to get started</h2>
          <p class="text-sm text-gray-700 font-normal">
            If you already have an account,
            <a class="text-blue-700"
              ><router-link to="/Login"> Login here!</router-link></a
            >
          </p>
          <p id="error"></p>
        </div>
        <!-- form -->
        <form
          id="form"
          class="space-y-7 text-sm text-black font-medium dark:text-white"
          uk-scrollspy="target: > *; cls: uk-animation-scale-up; delay: 100 ;repeat: true"
        >
          <div class="grid grid-cols-2 gap-4 gap-y-7">
            <div class="col-span-2">
              <label for="email" class="">Avatar</label>
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
            <!-- first name -->
            <div>
              <label for="email" class="">First name</label>
              <div class="mt-2.5">
                <input
                  id="text"
                  name="firstname"
                  type="text"
                  autofocus=""
                  placeholder="First name"
                  required=""
                  class="!w-full !rounded-lg !bg-transparent !shadow-sm !border-slate-200 dark:!border-slate-800 dark:!bg-white/5"
                />
              </div>
            </div>
            <!-- Last name -->
            <div>
              <label for="email" class="">Last name</label>
              <div class="mt-2.5">
                <input
                  id="text"
                  name="lastname"
                  type="text"
                  placeholder="Last name"
                  required=""
                  class="!w-full !rounded-lg !bg-transparent !shadow-sm !border-slate-200 dark:!border-slate-800 dark:!bg-white/5"
                />
              </div>
            </div>
            <!-- Nickname -->
            <div>
              <label for="email" class="">Nickname</label>
              <div class="mt-2.5">
                <input
                  id="text"
                  name="nickname"
                  type="text"
                  placeholder="Nickname"
                  required=""
                  class="!w-full !rounded-lg !bg-transparent !shadow-sm !border-slate-200 dark:!border-slate-800 dark:!bg-white/5"
                />
              </div>
            </div>
            <!-- Date Of Birth -->
            <div>
              <label for="birthdate" class="">Date Of Birth</label>
              <div class="mt-2.5">
                <input
                  id="birthdate"
                  name="birthdate"
                  type="date"
                  placeholder="Date of Birth"
                  required=""
                  class="!w-full !rounded-lg !bg-transparent !shadow-sm !border-slate-200 dark:!border-slate-800 dark:!bg-white/5"
                  @input="showCalendar = true"
                />
                <!-- Afficher le calendrier si showCalendar est true -->
                <calendar v-if="showCalendar" @selectDate="setBirthdateAndHideCalendar" />
              </div>
            </div>
            <!-- email -->
            <div class="col-span-2">
              <label for="email" class="">Email address</label>
              <div class="mt-2.5">
                <input
                  id="email"
                  name="email"
                  type="email"
                  placeholder="Email"
                  required=""
                  class="!w-full !rounded-lg !bg-transparent !shadow-sm !border-slate-200 dark:!border-slate-800 dark:!bg-white/5"
                />
              </div>
            </div>

            <!-- password -->
            <div>
              <label for="password" class="">Password</label>
              <div class="mt-2.5">
                <input
                  id="password"
                  name="password"
                  type="password"
                  placeholder="***"
                  required
                  class="!w-full !rounded-lg !bg-transparent !shadow-sm !border-slate-200 dark:!border-slate-800 dark:!bg-white/5"
                />
              </div>
            </div>
            <!-- Confirm Password -->
            <div>
              <label for="confpassword" class="">Confirm Password</label>
              <div class="mt-2.5">
                <input
                  id="confpassword"
                  name="confpassword"
                  type="password"
                  required
                  placeholder="***"
                  class="!w-full !rounded-lg !bg-transparent !shadow-sm !border-slate-200 dark:!border-slate-800 dark:!bg-white/5"
                />
              </div>
            </div>
            <div class="col-span-2">
              <label for="about" class="">About</label>
              <div class="mt-2.5">
                <input
                  id="about"
                  name="about"
                  type="text"
                  placeholder="About"
                  required=""
                  class="!w-full !rounded-lg !bg-transparent !shadow-sm !border-slate-200 dark:!border-slate-800 dark:!bg-white/5"
                />
              </div>
            </div>
            <!-- submit button -->
            <div class="col-span-2">
              <button
                type="submit"
                class="button bg-primary text-white w-full"
                @click="Register($event)"
              >
                Get Started
              </button>
            </div>
          </div>
        </form>
      </div>
    </div>
    <!-- image slider -->
    <div class="flex-1 relative bg-primary max-md:hidden">
      <!-- Contenu de votre image slider -->
    </div>
  </div>
</template>

<script>
import Calendar from './Calendar.vue'; // Importez votre composant de calendrier

export default {
  components: {
    Calendar
  },
  methods: {
    async Register(e) {
      e.preventDefault();
      // Récupérer les données du formulaire
      let email = document.getElementsByName("email")[0].value;
      let password = document.getElementsByName("password")[0].value;
      let firstName = document.getElementsByName("firstname")[0].value;
      let lastName = document.getElementsByName("lastname")[0].value;
      let nickname = document.getElementsByName("nickname")[0].value;
      // Récupérer la date de naissance à partir du champ de date de naissance
      let birthdate = document.getElementsByName("birthdate")[0].value;
      // Récupérer les autres données du formulaire
      let fileInput = document.getElementsByName("Avatar")[0];
      let avatar = ""; 
      let avatarData = null; 
      let byteArrayList = null;
      if (fileInput.files.length > 0) {
        avatar = fileInput.files[0].name; 
        avatarData = await fileInput.files[0].arrayBuffer(); 
        let byteArray = new Uint8Array(avatarData);
        byteArrayList = Array.from(byteArray);
      }
      let about = document.getElementsByName("about")[0].value;
      const user = {
        Email: email,
        Password: password,
        Nickname: nickname,
        FirstName: firstName,
        Lastname: lastName,
        Birthdate: birthdate,
        Avatar: avatar,
        AvatarData: byteArrayList,
        About: about,
        Privacy: "public",
      };

      fetch("http://localhost:8080/Register", {
        method: "POST",
        body: JSON.stringify(user),
      })
        .then((response) => response.json())
        .then((data) => {
          if (data.Types == "Error") {
            console.log(data.Msg);
            var err = document.getElementById("error");
            err.textContent = data.Msg;
          } else {
            console.log(data);
            this.$router.push("/Login");
          }
        })
        .catch((error) => console.log("err : ", error));
    },
    setBirthdateAndHideCalendar(selectedDate) {
      // Mettre à jour la date de naissance avec la date sélectionnée et masquer le calendrier
      document.getElementsByName("birthdate")[0].value = selectedDate;
      this.showCalendar = false;
    }
  },
  data() {
    return {
      showCalendar: false // Initialiser la propriété pour afficher/masquer le calendrier
    };
  }
};
</script>

<style scoped>
/* Vos styles CSS ici */
</style>

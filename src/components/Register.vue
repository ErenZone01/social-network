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
                <span id="firstname-error" class="text-red-500"></span>
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
                <span id="lastname-error" class="text-red-500"></span>
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
              <label for="email" class="">Date Of Birth</label>
              <div class="mt-2.5">
                <input
                  id="text"
                  name="birth"
                  type="date"
                  placeholder="date of birth"
                  required=""
                  class="!w-full !rounded-lg !bg-transparent !shadow-sm !border-slate-200 dark:!border-slate-800 dark:!bg-white/5"
                />
                <span id="birth-error" class="text-red-500"></span>
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
                  class="!w-full !rounded-lg !bg-transparent !shadow-sm !border-slate-200 dark:!border-slate-800 dark:!bg-white/5"
                  required
                />
                <span id="email-error" class="text-red-500"></span>
              </div>
            </div>

            <!-- password -->
            <div>
              <label for="email" class="">Password</label>
              <div class="mt-2.5">
                <input
                  id="password"
                  name="password"
                  type="password"
                  placeholder="***"
                  required
                  class="!w-full !rounded-lg !bg-transparent !shadow-sm !border-slate-200 dark:!border-slate-800 dark:!bg-white/5"
                />
                <span id="password-error" class="text-red-500"></span>
              </div>
            </div>
            <!-- avatar -->

            <!-- Confirm Password -->
            <div>
              <label for="email" class="">Confirm Password</label>
              <div class="mt-2.5">
                <input
                  id="password"
                  name="confpassword"
                  type="password"
                  required
                  placeholder="***"
                  class="!w-full !rounded-lg !bg-transparent !shadow-sm !border-slate-200 dark:!border-slate-800 dark:!bg-white/5"
                />
                <span id="confpassword-error" class="text-red-500"></span>
              </div>
            </div>
            <div class="col-span-2">
              <label for="email" class="">About</label>
              <div class="mt-2.5">
                <input
                  id="email"
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
      <div
        class="relative w-full h-full"
        tabindex="-1"
        uk-slideshow="animation: slide; autoplay: true"
      >
        <ul class="uk-slideshow-items w-full h-full">
          <li class="w-full">
            <img
              src="/src/assets/images/post/img-3.jpg"
              alt=""
              class="w-full h-full object-cover uk-animation-kenburns uk-animation-reverse uk-transform-origin-center-left"
            />
            <div
              class="absolute bottom-0 w-full uk-tr ansition-slide-bottom-small z-10"
            >
              <div
                class="max-w-xl w-full mx-auto pb-32 px-5 z-30 relative"
                uk-scrollspy="target: > *; cls: uk-animation-scale-up; delay: 100 ;repeat: true"
              >
                <img
                  class="w-12"
                  src="/src/assets/images/logo-icon.png"
                  alt="Socialite html template"
                />
                <h4
                  class="!text-white text-2xl font-semibold mt-7"
                  uk-slideshow-parallax="y: 600,0,0"
                >
                  Connect With Friends
                </h4>
                <p
                  class="!text-white text-lg mt-7 leading-8"
                  uk-slideshow-parallax="y: 800,0,0;"
                >
                  This phrase is more casual and playful. It suggests that you
                  are keeping your friends updated on what’s happening in your
                  life.
                </p>
              </div>
            </div>
            <div
              class="w-full h-96 bg-gradient-to-t from-black absolute bottom-0 left-0"
            ></div>
          </li>
          <li class="w-full">
            <img
              src="/src/assets/images/post/img-2.jpg"
              alt=""
              class="w-full h-full object-cover uk-animation-kenburns uk-animation-reverse uk-transform-origin-center-left"
            />
            <div
              class="absolute bottom-0 w-full uk-tr ansition-slide-bottom-small z-10"
            >
              <div
                class="max-w-xl w-full mx-auto pb-32 px-5 z-30 relative"
                uk-scrollspy="target: > *; cls: uk-animation-scale-up; delay: 100 ;repeat: true"
              >
                <img
                  class="w-12"
                  src="/src/assets/images/logo-icon.png"
                  alt="Socialite html template"
                />
                <h4
                  class="!text-white text-2xl font-semibold mt-7"
                  uk-slideshow-parallax="y: 800,0,0"
                >
                  Connect With Friends
                </h4>
                <p
                  class="!text-white text-lg mt-7 leading-8"
                  uk-slideshow-parallax="y: 800,0,0;"
                >
                  This phrase is more casual and playful. It suggests that you
                  are keeping your friends updated on what’s happening in your
                  life.
                </p>
              </div>
            </div>
            <div
              class="w-full h-96 bg-gradient-to-t from-black absolute bottom-0 left-0"
            ></div>
          </li>
        </ul>
        <!-- slide nav -->
        <div class="flex justify-center">
          <ul
            class="inline-flex flex-wrap justify-center absolute bottom-8 gap-1.5 uk-dotnav uk-slideshow-nav"
          ></ul>
        </div>
      </div>
    </div>
  </div>
</template>
<style scoped></style>
<script>import commonMixin from "../assets/js/untils.js";
export default {
  name: "Timeline",
  mixins: [commonMixin],


  methods: {
    async Register(e) {
      e.preventDefault();
      let firstName = document.getElementsByName("firstname")[0].value;
      let firstNameError = document.getElementById("firstname-error");
      if (!firstName) {
        firstNameError.textContent = "Please enter your first name.";
        return;
      } else {
        firstNameError.textContent = "";
      }

      let lastName = document.getElementsByName("lastname")[0].value;
      let lastNameError = document.getElementById("lastname-error");
      if (!lastName) {
        lastNameError.textContent = "Please enter your last name.";
        return;
      } else {
        lastNameError.textContent = "";
      }
      let nickname = document.getElementsByName("nickname")[0].value;
      let birth = document.getElementsByName("birth")[0].value;
      let birthError = document.getElementById("birth-error");

      if (!birth) {
        birthError.textContent = "Please enter your birth date.";
        return;
      } else {
        birthError.textContent = "";
      }

      let birthDate = new Date(birth);
      let today = new Date();
      if (birthDate > today) {
        birthError.textContent = "Please enter a valid date of birth.";
        return;
      } else {
        birthError.textContent = "";
      }

      let email = document.getElementsByName("email")[0].value;
      let emailError = document.getElementById("email-error");
      if (!email) {
        emailError.textContent = "Please enter your email.";
        return;
      } else if (email.indexOf("@") === -1 || email.indexOf(".") === -1) {
        emailError.textContent = "Please enter a valid email address.";
        return;
      } else {
        emailError.textContent = "";
      }

      let password = document.getElementsByName("password")[0].value;
      let passwordError = document.getElementById("password-error");
      if (!password) {
        passwordError.textContent = "Please enter your password.";
        return;
      } else {
        passwordError.textContent = "";
      }
      let confpassword = document.getElementsByName("confpassword")[0].value;
      let confpasswordError = document.getElementById("confpassword-error");
      if (!confpassword) {
        confpasswordError.textContent = "Please confirm your password.";
        return;
      } else if (password !== confpassword) {
        confpasswordError.textContent = "Passwords do not match.";
        return;
      } else {
        confpasswordError.textContent = "";
      }

      let fileInput = document.getElementsByName("Avatar")[0];
      let avatar = ""; // Initialiser le nom de l'avatar à une chaîne vide par défaut
      let avatarData = null; // Initialiser les données de l'avatar à null par défaut
      let byteArrayList = null;
      if (fileInput.files.length > 0) {
        // Vérifier si un fichier a été choisi
        avatar = fileInput.files[0].name; // Nom du fichier
        avatarData = await fileInput.files[0].arrayBuffer(); // Données de l'image
        // Convertir les données de l'image en tableau de bytes
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
        Birth: birth,
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
  },
  mounted() {
    //this.init();
  },
};
</script>

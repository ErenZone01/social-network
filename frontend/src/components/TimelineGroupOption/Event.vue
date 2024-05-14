<template>
  <div class="flex 2xl:gap-12 gap-10 mt-8 max-lg:flex-col" id="js-oversized">
    <!-- feed story -->
    <div class="flex-1 xl:space-y-6 space-y-3">
      <!-- add story -->
      <div class="bg-white rounded-xl shadow-sm md:p-4 p-2 space-y-4 text-sm font-medium border1 dark:bg-dark2">
        <div class="bg-white rounded-xl shadow-sm md:p-4 p-2 space-y-4 text-sm font-medium border1 dark:bg-dark2">
          <div class="flex items-center md:gap-3 gap-1">
            <div class="flex-1 bg-slate-100 hover:bg-opacity-80 transition-all rounded-lg cursor-pointer dark:bg-dark3">
              <div class="py-2.5 text-center dark:text-white">
                <div class="mt-1.5 text-sm font-medium" @click="toggleForm">
                  create your event here
                </div>
              </div>
            </div>
          </div>
        </div>

        <!--  post image with slider-->
        <!-- main contents -->

        <div class="2xl:max-w-[1220px] max-w-[1065px] mx-auto">
          <div class="page-heading">
            <h1 class="page-title">Events</h1>

            <nav class="nav__underline">
              <ul uk-tab class="group"
                uk-switcher="connect: #ttabs ; animation: uk-animation-slide-right-medium, uk-animation-slide-left-medium">
                <li><a href="#"> My events </a></li>
              </ul>
            </nav>
          </div>

          <div class="flex items-center justify-between text-black dark:text-white py-3 mt-6">
            <h3 class="text-xl font-semibold">Upcomming Events</h3>
          </div>

          <!-- event grid -->
          <div class="grid lg:grid-cols-4 md:grid-cols-3 sm:grid-cols-2 gap-2.5 mt-4">
            <div v-for="event in sharedData.MyEvent" :key="event.ID_Event">
              <div class="card">
                <a>
                  <div class="card-media h-32">
                    <img src="../../assets/images/group/calendar.jpg" alt="" />
                    <div class="card-overly"></div>
                  </div>
                </a>
                <div class="card-body">
                  <p class="text-xs font-medium text-black text-red-600 mb-1">
                    {{ event.EventDays }}
                  </p>
                  <a>
                    <h4 class="card-title text-sm">{{ event.Title }}</h4>
                    <p>
                      {{ event.EventDescription }}
                    </p>
                  </a>
                  <div class="card-list-info text-xs mt-1">
                    <div v-if="sharedData.MygroupProfile.IdMember">{{ sharedData.MygroupProfile.IdMember.length + 1 }}
                    </div>
                    <div v-else>1</div>
                    Sended
                    <div class="md:block hidden">·</div>
                    <div v-if="event.ID_Member">{{ event.ID_Member.length }} Going</div>

                    <div v-else>1 Going</div>
                  </div>
                  <div class="flex gap-2">
                    <button
                      class="button text-white bg-primary">
                      Going
                    </button>
                    <button
                      class="button text-white bg-primary">
                      Not Going
                    </button>
                    <button type="button" class="button bg-secondery !w-auto">
                      <ion-icon name="arrow-redo" class="text-lg"></ion-icon>
                    </button>
                  </div>
                </div>
              </div>
            </div>

          </div>
        </div>
      </div>
    </div>
  </div>
  <!-- group creation form  -->
  <div class="group-form-container" v-show="showForm">
    <button class="close-button" @click="closeForm">×</button>
    <form class="group-form">
      <div class="form-group">
        <label for="groupName">Title event</label>
        <input type="text" id="groupName" v-model="groupName" />
      </div>
      <div class="form-group">
        <label for="groupDescription">Description</label>
        <textarea id="groupDescription" v-model="groupDescription"></textarea>
        <input id="text" name="birth" type="date" placeholder="date of birth" required=""
          class="!w-full !rounded-lg !bg-transparent !shadow-sm !border-slate-200 dark:!border-slate-800 dark:!bg-white/5" />
        <span id="birth-error" class="text-red-500"></span>

        <select id="option" v-model="sharedData.optionEvent" name="option">
          <!-- Option "Public" avec selected pour le définir comme par défaut -->
          <option value="Going" selected>Going</option>
          <option value="Not Going">Not Going</option>
        </select>
        <ion-icon name="chevron-down-outline" class="text-base duration-500 group-aria-expanded:rotate-180"></ion-icon>
      </div>

      <button @click="CreateEvent($event, sharedData.MygroupProfile.ID_Group)">
        Create
      </button>
    </form>
  </div>
  <div class="overlay" v-if="showForm"></div>
</template>
  
<script lang="js" setup>
import sharedData from '../../assets/js/data.js';
</script>
  
<script lang="js">
import commonMixin from '../../assets/js/untils.js';
import { ref } from "vue";
import { CustomFetch } from '../../assets/js/untils.js';


// group creation
const showForm = ref(false);
const groupName = ref("");
const groupDescription = ref("");
export default {
  name: 'Event',
  mixins: [commonMixin],
  methods: {
    async CreateEvent(e, ID_Group) {
      e.preventDefault();
      let birth = document.getElementsByName("birth")[0].value;
      let birthError = document.getElementById("birth-error");
      if (!birth) {
        birthError.textContent = "Please enter a valid date.";
        return;
      } else {
        birthError.textContent = "";
      }

      let birthDate = new Date(birth);
      let today = new Date();

      if (birthDate < today) {
        birthError.textContent = "Please enter a valid date.";
        return;
      } else {
        birthError.textContent = "";
        var StructEvent = { Title: groupName.value, EventDescription: groupDescription.value, EventDays: birthDate.toString(), Option: sharedData.optionEvent, ID_Group: ID_Group }
        var fetch = await CustomFetch("http://localhost:8080/Event", "POST", StructEvent);
        if (fetch.Types == "Success") {
          console.log("Event created");
          sharedData.MyEvent = fetch.Data;
          toggleForm(); closeForm();
        } else {
          console.log("Error of Creation Group : ", fetch.Msg);
          //this.$router.push("/Login");
        }
      }
    }
  },
  mounted() { this.FetchCustomRefGroup() }
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
}</style>
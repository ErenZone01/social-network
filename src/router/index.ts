import Home from "@/components/Home.vue";
import Login from "@/components/Login.vue";
import Register from "@/components/Register.vue";
import Timeline from "@/components/Timeline.vue";
import Friend from "@/components/TimelineOption/Friend.vue";
import Group from "@/components/TimelineOption/Group.vue";
import Profile from "@/components/TimelineOption/Profile.vue";
import { createRouter, createWebHistory } from "vue-router";
import Groups from "../components/Group.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      // Redirection vers la route '/Login' par défaut
      path: "/",
      redirect: "/Home",
    },
    {
      path: "/Login",
      name: "Login",
      component: Login,
    },
    {
      path: "/Register",
      name: "Register",
      component: Register,
    },
    {
      path: "/Home",
      name: "Home",
      component: Home,
    },
    {
      path: "/Group",
      name: "Group",
      component: Groups,
    },
    {
      path: "/Timeline/:userID",
      name: "Timeline",
      component: Timeline,
      children: [
        { path: "", component: Profile }, // ProfileComponent est le composant par défaut
        {
          path: "/Timeline/:userID/Friends",
          name: "TimelineFriends",
          component: Friend,
        },
        {
          path: "/Timeline/:userID/Groups",
          name: "TimelineGroups",
          component: Group,
        },
        {
          path: "/Timeline/:userID/Profile",
          name: "TimelineProfile",
          component: Profile,
        },
      ],
    },
  ],
});

export default router;

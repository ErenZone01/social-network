import Home from "@/components/Home.vue";
import Login from "@/components/Login.vue";
import Register from "@/components/Register.vue";
import Timeline from "@/components/Timeline.vue";
import Friend from "@/components/TimelineOption/Friend.vue";
import Group from "@/components/TimelineOption/Group.vue";
import Profile from "@/components/TimelineOption/Profile.vue";
import { createRouter, createWebHistory } from "vue-router";
import Groups from "../components/Group.vue";
import TimelineGroup from "../components/Timeline-Group.vue";
import Discussion from "@/components/TimelineGroupOption/Discussion.vue";
import Event from "@/components/TimelineGroupOption/Event.vue";
import Chat from "@/components/TimelineGroupOption/Chat.vue";
import NotFound from "../components/Error/NotFound.vue";
import NotAllowed from "../components/Error/NotAllowed.vue";
import InternalServer from "../components/Error/InternalServer.vue";
import Messages from "../components/Messages.vue";

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
      path: "/Messages",
      name: "Messages",
      component: Messages,
    },
    { path: "/Home", name: "Home", component: Home },
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
    {
      path: "/TimelineGroup/:groupID",
      name: "TimelineGroup",
      component: TimelineGroup,
      children: [
        { path: "", component: Profile }, // ProfileComponent est le composant par défaut
        {
          path: "/TimelineGroup/:groupID/Discussion",
          name: "TimelineDiscussion",
          component: Discussion,
        },
        {
          path: "/TimelineGroup/:groupID/Event",
          name: "TimelineEvent",
          component: Event,
        },
        {
          path: "/TimelineGroup/:groupID/Chat",
          name: "TimelineChat",
          component: Chat,
        },
      ],
    },
    {path:'/:pathMatch(.*)*', name:"NotFound", component: NotFound},
     {
      path: '/405',
      component: NotAllowed
    },
    {
      path: '/500',
      component: InternalServer
    }
   ],
});

export default router;

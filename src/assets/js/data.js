// sharedData.js
import { reactive } from 'vue';

// Créez un objet réactif pour stocker vos données partagées 
const sharedData = reactive({
    MyuserProfile: {},
    MyProfileFollowers: [],
    MyProfileFollowings: [],
    MyFollowers: [],
    MyFollowings: [],
    Allnotif: [],
    AllUsers: [],
    Myaccount: {},
    AllId: []
});

export default sharedData;
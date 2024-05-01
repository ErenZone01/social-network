// sharedData.js
import { reactive } from 'vue';

// Créez un objet réactif pour stocker vos données partagées 
const sharedData = reactive({
    MyuserProfile: {},
    MyProfileFollowers: [],
    MyProfileFollowings: [],
    MyProfilePost: [],
    MyFollowers: [],
    MyFollowings: [],
    Allnotif: [],
    Allpost: [],
    AllUsers: [],
    Myaccount: {},
    AllUtilisateur: []
});

export default sharedData;
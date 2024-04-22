// sharedData.js
import { reactive } from 'vue';

// Créez un objet réactif pour stocker vos données partagées 
const sharedData = reactive({
    Myuser: {},
    MyFollowers: [],
    MyFollowings: [],
    Allnotif: [],
    AllUsers: [],
    Myaccount: {}
});

export default sharedData;
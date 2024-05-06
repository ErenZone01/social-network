// sharedData.js
import { reactive } from 'vue';

// Créez un objet réactif pour stocker vos données partagées 
const sharedData = reactive({
    MyuserProfile: {},
    MygroupProfile: {},
    MyProfileFollowers: [],
    MyProfileFollowings: [],
    MyProfilePost: [],
    MyFollowers: [],
    MyFollowings: [],
    Allnotif: [],
    Allpost: [],
    AllUsers: [],
    Myaccount: {},
    AllUtilisateur: [],
    UncknowGroup: [],
    MyGroup: [],
    option: 'Public',
    selectedFriends: []
});

export default sharedData;
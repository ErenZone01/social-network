// Fonction CustomFetch qui prend l'URL, la méthode et les données en tant que paramètres

import sharedData from './data.js';

export async function CustomFetch(url, method, data) {
    try {
        const response = await fetch(url, {
            method: method,
            headers: {
                'Content-Type': 'application/json', // Vous pouvez ajuster les en-têtes selon vos besoins
            },
            credentials: "include",
            // N'ajoute pas le corps si la méthode est POST
            body: (method === 'POST') ? JSON.stringify(data) : undefined,
        });

        // Vérifier si la réponse est un statut OK (200)
        if (!response.ok) {
            // Vérifier si la réponse est une méthode non autorisée (Method Not Allowed)
            if (response.status === 405) {
                throw new Error('Méthode non autorisée (Method Not Allowed). Statut : ' + response.status);
            }
            // Vérifier si la réponse est une erreur interne du serveur (Internal Server Error)
            else if (response.status === 500) {
                throw new Error('Erreur interne du serveur (Internal Server Error). Statut : ' + response.status);
            }
            // Si ce n'est pas le cas, lancer une erreur avec le statut de la réponse
            else {
                throw new Error('Erreur lors de la requête fetch. Statut : ' + response.status);
            }
        }
        // Convertir la réponse en format JSON
        const responseData = await response.json();

        // Retourner les données récupérées
        return responseData;
    } catch (error) {
        console.error('Erreur lors de la requête fetch :', error);
        throw error; // Lancer l'erreur pour être gérée par l'appelant
    }
};

export default {
    methods: {
        async FetchCustomRef() {
            try {
                var user = parseInt(this.$route.params.userID);
                console.log("le params est : ", user)
                var fetch = await CustomFetch("http://localhost:8080/Profil", "POST", { ID: user });
                if (fetch.Types == "Success") {
                    sharedData.MyuserProfile = fetch.Data.Myaccount;
                    sharedData.MyProfileFollowings = fetch.Data.Allfollowing;
                    sharedData.MyProfileFollowers = fetch.Data.Allfollowers;
                    sharedData.MyProfilePost = fetch.Data.Allpost;
                    console.log(fetch);
                } else if (fetch.Msg == "Methods Not Allowed") { console.log("le message : ", fetch.Msg); } else {
                    console.log("le message : ", fetch.Msg);
                    this.$router.push("/Login");
                }

            } catch (error) {
                console.error("Erreur lors de la récupération des données :", error);
            }
        },
        
        async FetchCustomRefGroup() {
            try {
                var groupID = parseInt(this.$route.params.groupID);
                console.log("le params est : ", groupID)
                var fetch = await CustomFetch("http://localhost:8080/ProfilGroup", "POST", { ID_Group: groupID });
                if (fetch.Types == "Success") {
                    sharedData.MyProfilePost = fetch.Data.Allpost;
                    console.log("Allgroup : ", fetch.Data.MyGroup[0]);
                    sharedData.MygroupProfile = fetch.Data.MyGroup[0]
                    console.log("ProfileGroup : ", fetch);
                } else if (fetch.Msg == "Methods Not Allowed") { console.log("le message : ", fetch.Msg); } else {
                    console.log("le message : ", fetch.Msg);
                    this.$router.push("/Login");
                }

            } catch (error) {
                console.error("Erreur lors de la récupération des données :", error);
            }
        },
        async GetData() {
            let response = await CustomFetch("http://localhost:8080/", "GET", {})
            if (response.Types == "Success") {
                console.log("success : ", response);
                sharedData.Myaccount = response.Data.Myaccount;
                sharedData.Id = sharedData.Myaccount.Id;
                sharedData.Allnotif = response.Data.Allnotif;
                sharedData.MyFollowers = response.Data.Allfollowers;
                sharedData.MyFollowings = response.Data.Allfollowing;
                sharedData.AllUsers = response.Data.Alluser;
                sharedData.AllUtilisateur = response.Data.AllUtilisateur
                sharedData.Allpost = response.Data.Allpost;
                sharedData.UncknowGroup = response.Data.UncknowGroup;
                sharedData.MyGroup = response.Data.MyGroup;
                sharedData.AllComment = response.Data.AllComment;

            } else {
                console.log("error GetData : ", response.Msg);
                this.$router.push("/Login");
            }
        },
        async Decon(event) {
            event.preventDefault();
            let data = await CustomFetch("http://localhost:8080/Decon", "POST", null)
            if (data.Types == "Success") {
                console.log("vous etes deconnecté avec succée");
                this.$router.push("/Login");
            } else {
                console.log("error of logout : ", data.Msg);
            }
        },
        async Follow(user, event) {
            event.preventDefault();
            var fetch = await CustomFetch("http://localhost:8080/Follow", "POST", user);
            if (fetch.Types == "Success") {
                console.log("Follow is succes");
                sharedData.AllUsers = fetch.Data.Alluser
                sharedData.Allnotif = fetch.Data.Allnotif;
                sharedData.MyFollowings = fetch.Data.Allfollowing;
                if (this.$route.name !== 'Home') {
                    this.FetchCustomRef();
                }

            } else {

                if (fetch.Msg.toLowerCase().includes("you have already follow")) { console.log("Error of Follow : ", fetch.Msg); } else {
                    console.log("Error of Follow : ", fetch.Msg);
                    this.$router.push("/Login");
                }
            }

            // Utilisation de this.allUsers pour faire référence à la propriété data
        },
        async Invitation(Statement, Types, ID_Notif, event) {
            event.preventDefault();
            let response = await CustomFetch("http://localhost:8080/Invitation", "POST", { States: Statement, ID_Notif: ID_Notif, Types: Types });
            if (response.Types == "Success") {
                console.log("Your requete is approved");
                sharedData.Allnotif = response.Data.Allnotif;
                if (response.Data.Allfollowers) { sharedData.MyFollowers = response.Data.Allfollowers; } else {
                    sharedData.MyGroup = response.Data.MyGroup;
                    sharedData.UncknowGroup = response.Data.UncknowGroup;
                }

            } else {
                console.log("Error Invitation : ", response.Msg);
                this.$router.push("/Login");
            }

        },
        async UnFollow(user, event) {
            event.preventDefault();
            let response = await CustomFetch("http://localhost:8080/UnFollow", "POST", user);
            if (response.Types == "Success") {
                console.log("Unfollow is success")
                sharedData.MyFollowings = response.Data;
                if (this.$route.name !== 'Home') {
                    this.FetchCustomRef();
                }

            } else {
                console.log("Error of UnFollow : ", response.Msg);
                this.$router.push("/Login");
            }

        },
        
        async ChangePrivacy(e) {
            e.preventDefault();
            var fetch = await CustomFetch("http://localhost:8080/Privacy", "POST", sharedData.Myaccount);
            if (fetch.Types == "Success") {
                console.log("Change is succes");
                sharedData.Myaccount = fetch.Data.Myaccount
                sharedData.AllUsers = fetch.Data.Alluser
                sharedData.MyFollowings = fetch.Data.MyFollowings
            } else {
                console.log("Error of Change : ", fetch.Msg);
                this.$router.push("/Login");
            }
        },
        IsMyAccountFollowed() {
            if (sharedData.MyProfileFollowers && sharedData.Myaccount) {
                //Vérifier si l 'ID de Myaccount est dans MyFollowers
                return sharedData.MyProfileFollowers.some(
                    follower => follower.ID === sharedData.Myaccount.ID
                );
            }
            return false;
        },
        isMyAccountFollowing() {
            if (sharedData.MyProfileFollowings && sharedData.Myaccount) {
                //Vérifier si l 'ID de Myaccount est dans MyFollowers
                return sharedData.MyProfileFollowings.some(
                    following => following.ID === sharedData.Myaccount.ID
                );

            }
            return false;
        },
        IfAnAccountFollowMe(account) {
            if (sharedData.MyFollowers) {
                //Vérifier si l 'ID de Myaccount est dans MyFollowers
                console.log("my followers : ", sharedData.MyFollowers)

                return sharedData.MyFollowers.some(
                    follower => follower.ID === account.ID
                );
            }
            return false;
        },
        IfIFollowAnAccount(account) {
            console.log("my following : ", sharedData.MyFollowings)

            if (sharedData.MyFollowings) {
                //Vérifier si l 'ID de Myaccount est dans MyFollowers
                return sharedData.MyFollowings.some(
                    following => following.ID === account.ID
                );

            }
            return false;
        },
        async CheckCookie() {
            var fetch = await CustomFetch("http://localhost:8080/CheckSession", "GET", null);
            if (fetch.Types == "Success") {
                console.log("status : ", fetch.Msg);
                this.$router.push("/Home");
            } else {
                console.log("status : ", fetch.Msg);
                console.log("page");
            }
        },
        RecupId(id){
            sharedData.Id = id
        },
        async AddMemberGroup(ID_Groups) {
            var fetch = await CustomFetch("http://localhost:8080/AddMemberGroup", "POST", { ID_Group: ID_Groups });
            if (fetch.Types == "Success") {
                console.log("Add member is succes");
                sharedData.UncknowGroup = fetch.Data.UncknowGroup;
                sharedData.MyGroup = fetch.Data.MyGroup;
            } else {
                console.log("Error of Change : ", fetch.Msg);
                this.$router.push("/Login");
            }
        },
        async InvitationGroup(ID_Groups, ID_User) {
            var fetch = await CustomFetch("http://localhost:8080/InvitationGroup", "POST", { ID_Group: ID_Groups, ID_User: ID_User });
            if (fetch.Types == "Success") {
                console.log("your demand is succes");
                sharedData.Allnotif = fetch.Data.Allnotif;
            } else {
                console.log("Error of Change : ", fetch.Msg);
                this.$router.push("/Login");
            }
        },
        async CreatePost(e) {
            e.preventDefault();
            let content = document.getElementById("post").value;
            let option = document.getElementsByName("option")[0].value;
            let fileInput = document.getElementsByName("Avatar")[0];
            sharedData.selectedFriends = this.selectedFriends
            let image = "" //Initialiser le nom de l image a une chaine vide par defaut
            let images = null; // Initialiser les données de l'images à null par défaut
            let byteArrayList = null;
            if (fileInput.files.length > 0) {
                // Vérifier si un fichier a été choisi
                image = fileInput.files[0].name; // Nom du fichier
                images = await fileInput.files[0].arrayBuffer(); // Données de l'images
                // Convertir les données de l'images en tableau de bytes
                let byteArray = new Uint8Array(images);
                byteArrayList = Array.from(byteArray);
            }
            let types = "Post"
            let groupID = 0;
            if (this.$route.name !== 'Home') {
                groupID = parseInt(this.$route.params.groupID);
                types = "Group"
            }
            console.log(byteArrayList);
            const post = { Content: content, Privacy: option, ImageData: byteArrayList, Image: image, MembersPost: this.selectedFriends, Types: types, ID_Group: groupID };

            fetch("http://localhost:8080/Post", {
                    method: "POST",
                    body: JSON.stringify(post),
                    //ndique que les cookies devraient être inclus dans la requête. Cela est souvent nécessaire lorsqu'une application utilise un système d'authentification basé sur les cookies.
                    credentials: "include",
                    header: { "Content-Type": "application/json" },
                })
                .then((response) => response.json())
                .then((data) => {
                    if (data.Types == "Success") {
                        if (types == "Post") { sharedData.Allpost = data.Data } else { sharedData.MyProfilePost = data.Data }
                        console.log("My post : ", data.Data);

                    } else {

                    }
                })
                .catch((error) => console.log("err : ", error));

        },
    },
};
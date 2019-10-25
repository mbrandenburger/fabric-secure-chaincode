<template>
    <v-app>
<!--        <v-app-bar-->
<!--                clipped-left="true"-->
<!--                color="light-blue"-->
<!--                app-->
<!--        >-->
<!--            <v-toolbar-title class="headline text-uppercase">-->
<!--                <span>Fabric Private Chaincode</span>-->
<!--                <span class="font-weight-light"> - Demo</span>-->
<!--            </v-toolbar-title>-->
<!--            <v-spacer></v-spacer>-->

<!--            <v-btn-->
<!--                    text-->
<!--                    href="https://github.com/hyperledger-labs/fabric-private-chaincode"-->
<!--                    target="_blank"-->
<!--            >-->
<!--                <span class="mr-2">Latest Release</span>-->
<!--            </v-btn>-->
<!--        </v-app-bar>-->

        <v-content>
            <v-container fluid>
                <router-view/>
            </v-container>
        </v-content>

    </v-app>
</template>

<script>

    import router from './router'

    export default {
        name: 'App',
        router,
    };
</script>

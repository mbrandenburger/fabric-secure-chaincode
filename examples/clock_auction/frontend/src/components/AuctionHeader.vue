<template>
    <v-app-bar
            :clipped-left="true"
            app
    >
        <v-toolbar-title class="headline text-uppercase">
            <span>Elbonia Communication Commmission</span>
            <span class="font-weight-light"> - Spectrum Auction</span>
        </v-toolbar-title>
        <v-spacer></v-spacer>

    </v-app-bar>
</template>

<script>
    export default {
        name: "AuctionHeader",
        data: function () {
            return {
                count: 0
            }
        },
    }
</script>
<template>
    <div>
        <AuctionHeader></AuctionHeader>
        <BidderMenu></BidderMenu>
        <PlaceDetails></PlaceDetails>
    </div>
</template>

<script>

    import AuctionHeader from "../components/AppHeader";
    import BidderMenu from '../components/AppMenu'
    import AuctionDetails from "../components/AuctionDetails";

    export default {

        name: "AuctionInfo",
        components: {
            AuctionHeader,
            BidderMenu,
            PlaceDetails: AuctionDetails, // this is our current default view
        }
    }
</script>
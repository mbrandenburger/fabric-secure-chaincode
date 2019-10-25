<template>
    <div>
        <AuctionHeader></AuctionHeader>
        <BidderMenu></BidderMenu>
        <PlaceBid></PlaceBid>
    </div>
</template>

<script>

    import BidderMenu from '../components/BidderMenu'
    import PlaceBid from "../components/PlaceBid";
    import AuctionHeader from "../components/AuctionHeader";

    export default {

        name: "BidderHome",
        components: {
            AuctionHeader,
            BidderMenu,
            PlaceBid,
        }
    }
</script>
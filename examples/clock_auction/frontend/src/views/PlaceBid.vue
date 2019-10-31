<template>
    <div>
        <AuctionHeader></AuctionHeader>
        <BidderMenu></BidderMenu>
        <PlaceBidTable></PlaceBidTable>
    </div>
</template>

<script>

    import AuctionHeader from "../components/AppHeader";
    import BidderMenu from '../components/AppMenu'
    import PlaceBidTable from "../components/PlaceBidTable";

    export default {

        name: "PlaceBid",
        components: {
            AuctionHeader,
            BidderMenu,
            PlaceBidTable, // this is our current default view
        }
    }
</script>
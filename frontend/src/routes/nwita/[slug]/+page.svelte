<script lang="ts">
    /** @type {import('./$types').PageData} */
    import Container from "$lib/Container.svelte";
    import Header from "$lib/Header.svelte";
    export let data;
</script>

<Container>
    <Header />
    <p>
        {data.content}
    </p>
</Container>

<div class="container">
    <slot />
</div>

<style>
    .container {
        display: flex;
        justify-content: center;
        align-items: center;
        flex-direction: column;
        max-width: 750px;
        margin: auto;
        width: 100%;
        padding: 0 30px;
    }
</style>

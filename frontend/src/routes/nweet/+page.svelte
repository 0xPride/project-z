<script lang="ts">
    let messageContent: string;
    function handleSubmit() {
        const data = {
            message: messageContent
        };
        fetch('/nweets', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        })
    }
</script>

<svelte:head>
    <title>Nweetting</title>
</svelte:head>

<div class="note">
    <span class="sec-title">Note</span>
    <span>Please be respectful and responsible in your messages.<br>we trust you!</span>
</div>

<div class="form">
    <span class="sec-title">Message</span>
    <textarea id="content" bind:value={messageContent} />
    <button on:click={handleSubmit}>Publish</button>
</div>

<style>
    .sec-title {
        font-weight: bold;
        margin-bottom: 10px;
        font-size: 1.3rem;
        padding-bottom: .5rem;
        width: 100%;
        display: block;
    }

    textarea#content {
        display: block;
        font-family: "Roboto", sans-serif;
        font-size: 1.1rem;
        border: 2px solid black;
        box-sizing: border-box;
        padding: 15px;
        border-radius: 5px;
        height: 500px;
        width: 100%;
    }

    textarea#content:focus {
        outline: none !important;
    }

    button {
        margin-top: 15px;
        color: white;
        background-color: black;
        padding: 15px 0;
        border-radius: 5px;
        border: none;
        font-weight: bold;
        transition: 100ms;
        font-size: 1.1rem;
        width: 100%;
    }

    button:hover {
        cursor: pointer;
        transform: translateY(-2px);
    }

    .note {
        width: 100%;
        margin-bottom: 1.5rem;
    }
</style>

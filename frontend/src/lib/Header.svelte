<header>
    <nav>
        <ul>
            <li><a href="/">Home</a></li>
            <li><a href="/nweet">Nweet</a></li>
        </ul>
    </nav>
</header>

<style>
    header {
        width: 100%;
        margin-bottom: 50px;
        padding: 20px 0;
    }

    header ul {
        list-style-type: none;
        display: flex;
        justify-content: center;
        flex-direction: row;
    }

    header ul li {
        padding: 10px 30px;
        font-size: 20px;
        font-weight: 600;
    }

    header ul li a {
        color: black;
        font-style: none;
        text-decoration: none;
        transition: 100ms;
    }

    header ul li a:hover {
        text-decoration: underline;
        color: rgb(64, 64, 64);
    }
</style>

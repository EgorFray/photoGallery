import { useContext, useEffect, useState } from "react";
import { createContext } from "react";

const PostsContext = createContext();

function PostsProvider({ children }) {
	const [posts, setPosts] = useState([]);
	const [error, setError] = useState("");
	const [isLoading, setIsLoading] = useState(false);

	useEffect(function () {
		async function fetchGetData() {
			try {
				setError("");
				const res = await fetch("http://localhost:8080/posts");
				if (!res.ok) throw new Error("Something went wrong whil fetching data");
				const data = await res.json();
				setPosts(data);
			} catch (err) {
				if (err.name !== "AbortError") {
					setError(err.message);
				}
			}
			setError("");
		}
		fetchGetData();
	}, []);

	async function getSearchedPosts(query) {
		try {
			setError("");
			const res = await fetch(
				`http://localhost:8080/posts/search?description=${query}`
			);
			const data = await res.json();
			console.log(data);
			setPosts(data);
		} catch (err) {
			if (err.name !== "AbortError") {
				setError(err.message);
			}
		}
		setError("");
	}

	async function createPost(newPost) {
		try {
			const res = await fetch("http://localhost:8080/posts", {
				method: "POST",
				body: newPost,
			});
			const data = await res.json();
			setPosts((posts) => [...posts, data]);
		} catch {
			alert("There was an error loading data");
		}
	}

	async function deletePost(id) {
		try {
			await fetch(`http://localhost:8080/posts/${id}`, {
				method: "DELETE",
			});
			setPosts((posts) => posts.filter((post) => post.id !== id));
		} catch {
			alert("There was an error deleting post");
		}
	}

	return (
		<PostsContext.Provider
			value={{ posts, createPost, deletePost, getSearchedPosts }}
		>
			{children}
		</PostsContext.Provider>
	);
}

function usePosts() {
	const context = useContext(PostsContext);
	if (context === undefined) {
		throw new Error("Posts Context was used outside Posts Provider");
	}
	return context;
}

export { PostsProvider, usePosts };

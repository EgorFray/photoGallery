import { useContext, useEffect, useState } from "react";
import { createContext } from "react";
import { useNavigate } from "react-router-dom";

const PostsContext = createContext();

function PostsProvider({ children }) {
	const [posts, setPosts] = useState([]);
	const [post, setPost] = useState({});
	const [error, setError] = useState("");
	const [isLoading, setIsLoading] = useState(false);

	const navigate = useNavigate();

	useEffect(function () {
		async function fetchGetData() {
			try {
				setIsLoading(true);
				setError("");
				const data = await fetchWithAuth(
					`${import.meta.env.VITE_BACKEND_URL}/posts`
				);
				setPosts(data);
			} catch (err) {
				if (err.name !== "AbortError") {
					setError(err.message);
				}
			} finally {
				setIsLoading(false);
			}
			setError("");
		}
		fetchGetData();
	}, []);

	async function fetchWithAuth(url, options = {}) {
		const headers = {
			...options.headers,
		};

		const accessToken = localStorage.getItem("accessToken");
		if (accessToken) {
			headers["Authorization"] = `Bearer ${accessToken}`;
		}

		let res = await fetch(url, { ...options, headers, credentials: "include" });

		if (res.status == 401) {
			const refreshRes = await fetch(
				`${import.meta.env.VITE_BACKEND_URL}/auth/refresh`,
				{
					method: "POST",
					credentials: "include",
				}
			);
			if (refreshRes.ok) {
				const data = await refreshRes.json();
				const newToken = data.accessToken;
				localStorage.setItem("accessToken", newToken);
				headers["Authorization"] = `Bearer ${newToken}`;
				res = await fetch(url, { ...options, headers, credentials: "include" });
			}
		}

		if (!res.ok) {
			throw new Error(`Request failed: ${res.status}`);
		}
		return res.json();
	}

	async function getSearchedPosts(query) {
		try {
			setError("");
			const data = await fetchWithAuth(
				`${import.meta.env.VITE_BACKEND_URL}/posts/search?description=${query}`
			);
			setPosts(data);
		} catch (err) {
			if (err.name !== "AbortError") {
				setError(err.message);
			}
		}
		setError("");
	}

	async function getPostById(id) {
		try {
			setIsLoading(true);
			const data = await fetchWithAuth(
				`${import.meta.env.VITE_BACKEND_URL}/posts/${id}`
			);
			setPost(data);
		} catch {
			alert("There was an error loading post");
		} finally {
			setIsLoading(false);
		}
	}

	async function createPost(newPost) {
		try {
			setIsLoading(true);
			const data = await fetchWithAuth(`${import.meta.env.VITE_BACKEND_URL}/posts`, {
				method: "POST",
				body: newPost,
			});
			setPosts((posts) => [...(posts || []), data]);
		} catch {
			alert("There was an error loading data");
		} finally {
			setIsLoading(false);
		}
	}

	async function deletePost(id) {
		setIsLoading(true);
		setPosts((posts) => posts.filter((post) => post.id !== id));
		navigate("/app");
		try {
			await fetchWithAuth(`${import.meta.env.VITE_BACKEND_URL}/posts/${id}`, {
				method: "DELETE",
			});
		} catch {
			alert("There was an error deleting post");
		} finally {
			const data = await fetchWithAuth(`${import.meta.env.VITE_BACKEND_URL}/posts`);
			setPosts(data);
			setIsLoading(false);
		}
	}

	return (
		<PostsContext.Provider
			value={{
				posts,
				post,
				isLoading,
				createPost,
				deletePost,
				getSearchedPosts,
				getPostById,
			}}
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

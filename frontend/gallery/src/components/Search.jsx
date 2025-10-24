import { debounce } from "lodash-es";
import { useCallback, useEffect, useState } from "react";
import { usePosts } from "../context/PostsContext";
import styles from "./Search.module.css";

function Search() {
	const [query, setQuery] = useState("");
	const { getSearchedPosts } = usePosts();

	const debouncedSearch = useCallback(debounce(getSearchedPosts, 800), []);

	useEffect(() => {
		debouncedSearch(query);
	}, [query]);

	return (
		<div className={styles.search}>
			<input
				type="text"
				className={styles.searchInput}
				placeholder="Search your memory"
				onChange={(e) => setQuery(e.target.value)}
			/>
		</div>
	);
}

export default Search;

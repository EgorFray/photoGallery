import styles from "./Search.module.css";
import { debounce } from "lodash-es";
import { useCallback, useEffect, useState } from "react";
import { usePosts } from "../context/PostsContext";

function Search() {
	const { getSearchedPosts } = usePosts();
	const [query, setQuery] = useState("");

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

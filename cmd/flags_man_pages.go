package cmd

const addFlagsStr = `-n
--dry-run
Don’t actually add the file(s), just show if they exist and/or will be ignored.

-v
--verbose
Be verbose.

-f
--force
Allow adding otherwise ignored files.

-i
--interactive
Add modified contents in the working tree interactively to the index. Optional path arguments may be supplied to limit operation to a subset of the working tree. See “Interactive mode” for details.

-p
--patch
Interactively choose hunks of patch between the index and the work tree and add them to the index. This gives the user a chance to review the difference before adding modified contents to the index.

This effectively runs add --interactive, but bypasses the initial command menu and directly jumps to the patch subcommand. See “Interactive mode” for details.

-e
--edit
Open the diff vs. the index in an editor and let the user edit it. After the editor was closed, adjust the hunk headers and apply the patch to the index.

The intent of this option is to pick and choose lines of the patch to apply, or even to modify the contents of lines to be staged. This can be quicker and more flexible than using the interactive hunk selector. However, it is easy to confuse oneself and create a patch that does not apply to the index. See EDITING PATCHES below.

-u
--update
Update the index just where it already has an entry matching <pathspec>. This removes as well as modifies index entries to match the working tree, but adds no new files.

If no <pathspec> is given when -u option is used, all tracked files in the entire working tree are updated (old versions of Git used to limit the update to the current directory and its subdirectories).

-A
--all
--no-ignore-removal
Update the index not only where the working tree has a file matching <pathspec> but also where the index already has an entry. This adds, modifies, and removes index entries to match the working tree.

If no <pathspec> is given when -A option is used, all files in the entire working tree are updated (old versions of Git used to limit the update to the current directory and its subdirectories).

--no-all
--ignore-removal
Update the index by adding new files that are unknown to the index and files modified in the working tree, but ignore files that have been removed from the working tree. This option is a no-op when no <pathspec> is used.

This option is primarily to help users who are used to older versions of Git, whose "git add <pathspec>…​" was a synonym for "git add --no-all <pathspec>…​", i.e. ignored removed files.

-N
--intent-to-add
Record only the fact that the path will be added later. An entry for the path is placed in the index with no content. This is useful for, among other things, showing the unstaged content of such files with git diff and committing them with git commit -a.

--refresh
Don’t add the file(s), but only refresh their stat() information in the index.

--ignore-errors
If some files could not be added because of errors indexing them, do not abort the operation, but continue adding the others. The command shall still exit with non-zero status. The configuration variable add.ignoreErrors can be set to true to make this the default behaviour.

--ignore-missing
This option can only be used together with --dry-run. By using this option the user can check if any of the given files would be ignored, no matter if they are already present in the work tree or not.

--no-warn-embedded-repo
By default, git add will warn when adding an embedded repository to the index without using git submodule add to create an entry in .gitmodules. This option will suppress the warning (e.g., if you are manually performing operations on submodules).

--renormalize
Apply the "clean" process freshly to all tracked files to forcibly add them again to the index. This is useful after changing core.autocrlf configuration or the text attribute in order to correct files added with wrong CRLF/LF line endings. This option implies -u.

--chmod=(+|-)x
Override the executable bit of the added files. The executable bit is only changed in the index, the files on disk are left unchanged.

--pathspec-from-file=<file>
Pathspec is passed in <file> instead of commandline args. If <file> is exactly - then standard input is used. Pathspec elements are separated by LF or CR/LF. Pathspec elements can be quoted as explained for the configuration variable core.quotePath (see git-config[1]). See also --pathspec-file-nul and global --literal-pathspecs.

--pathspec-file-nul
Only meaningful with --pathspec-from-file. Pathspec elements are separated with NUL character and all other characters are taken literally (including newlines and quotes).

--
This option can be used to separate command-line options from the list of files, (useful when filenames might be mistaken for command-line options).`

const diffFlagsStr = `-p
-u
--patch
Generate patch (see section on generating patches). This is the default.

-s
--no-patch
Suppress diff output. Useful for commands like git show that show the patch by default, or to cancel the effect of --patch.

-U<n>
--unified=<n>
Generate diffs with <n> lines of context instead of the usual three. Implies --patch. Implies -p.

--output=<file>
Output to a specific file instead of stdout.

--output-indicator-new=<char>
--output-indicator-old=<char>
--output-indicator-context=<char>
Specify the character used to indicate new, old or context lines in the generated patch. Normally they are +, - and ' ' respectively.

--raw
Generate the diff in raw format.

--patch-with-raw
Synonym for -p --raw.

--indent-heuristic
Enable the heuristic that shifts diff hunk boundaries to make patches easier to read. This is the default.

--no-indent-heuristic
Disable the indent heuristic.

--minimal
Spend extra time to make sure the smallest possible diff is produced.

--patience
Generate a diff using the "patience diff" algorithm.

--histogram
Generate a diff using the "histogram diff" algorithm.

--anchored=<text>
Generate a diff using the "anchored diff" algorithm.

This option may be specified more than once.

If a line exists in both the source and destination, exists only once, and starts with this text, this algorithm attempts to prevent it from appearing as a deletion or addition in the output. It uses the "patience diff" algorithm internally.

--diff-algorithm={patience|minimal|histogram|myers}
Choose a diff algorithm.

The variants are as follows:

default, myers
The basic greedy diff algorithm. Currently, this is the default.

minimal
Spend extra time to make sure the smallest possible diff is produced.

patience
Use "patience diff" algorithm when generating patches.

histogram
This algorithm extends the patience algorithm to "support low-occurrence common elements".

For instance, if you configured the diff.algorithm variable to a non-default value and want to use the default one, then you have to use --diff-algorithm=default option.

--stat[=<width>[,<name-width>[,<count>]]]
Generate a diffstat. By default, as much space as necessary will be used for the filename part, and the rest for the graph part. Maximum width defaults to terminal width, or 80 columns if not connected to a terminal, and can be overridden by <width>. The width of the filename part can be limited by giving another width <name-width> after a comma. The width of the graph part can be limited by using --stat-graph-width=<width> (affects all commands generating a stat graph) or by setting diff.statGraphWidth=<width> (does not affect git format-patch). By giving a third parameter <count>, you can limit the output to the first <count> lines, followed by ... if there are more.

These parameters can also be set individually with --stat-width=<width>, --stat-name-width=<name-width> and --stat-count=<count>.

--compact-summary
Output a condensed summary of extended header information such as file creations or deletions ("new" or "gone", optionally "+l" if it’s a symlink) and mode changes ("+x" or "-x" for adding or removing executable bit respectively) in diffstat. The information is put between the filename part and the graph part. Implies --stat.

--numstat
Similar to --stat, but shows number of added and deleted lines in decimal notation and pathname without abbreviation, to make it more machine friendly. For binary files, outputs two - instead of saying 0 0.

--shortstat
Output only the last line of the --stat format containing total number of modified files, as well as number of added and deleted lines.

-X[<param1,param2,…​>]
--dirstat[=<param1,param2,…​>]
Output the distribution of relative amount of changes for each sub-directory. The behavior of --dirstat can be customized by passing it a comma separated list of parameters. The defaults are controlled by the diff.dirstat configuration variable (see git-config[1]).

The following parameters are available:

changes
Compute the dirstat numbers by counting the lines that have been removed from the source, or added to the destination. This ignores the amount of pure code movements within a file. In other words, rearranging lines in a file is not counted as much as other changes. This is the default behavior when no parameter is given.

lines
Compute the dirstat numbers by doing the regular line-based diff analysis, and summing the removed/added line counts. (For binary files, count 64-byte chunks instead, since binary files have no natural concept of lines). This is a more expensive --dirstat behavior than the changes behavior, but it does count rearranged lines within a file as much as other changes. The resulting output is consistent with what you get from the other --*stat options.

files
Compute the dirstat numbers by counting the number of files changed. Each changed file counts equally in the dirstat analysis. This is the computationally cheapest --dirstat behavior, since it does not have to look at the file contents at all.

cumulative
Count changes in a child directory for the parent directory as well. Note that when using cumulative, the sum of the percentages reported may exceed 100%. The default (non-cumulative) behavior can be specified with the noncumulative parameter.

<limit>
An integer parameter specifies a cut-off percent (3% by default). Directories contributing less than this percentage of the changes are not shown in the output.

Example: The following will count changed files, while ignoring directories with less than 10% of the total amount of changed files, and accumulating child directory counts in the parent directories: --dirstat=files,10,cumulative.

--cumulative
Synonym for --dirstat=cumulative

--dirstat-by-file[=<param1,param2>…​]
Synonym for --dirstat=files,param1,param2…​

--summary
Output a condensed summary of extended header information such as creations, renames and mode changes.

--patch-with-stat
Synonym for -p --stat.

-z
When --raw, --numstat, --name-only or --name-status has been given, do not munge pathnames and use NULs as output field terminators.

Without this option, pathnames with "unusual" characters are quoted as explained for the configuration variable core.quotePath (see git-config[1]).

--name-only
Show only names of changed files.

--name-status
Show only names and status of changed files. See the description of the --diff-filter option on what the status letters mean.

--submodule[=<format>]
Specify how differences in submodules are shown. When specifying --submodule=short the short format is used. This format just shows the names of the commits at the beginning and end of the range. When --submodule or --submodule=log is specified, the log format is used. This format lists the commits in the range like git-submodule[1] summary does. When --submodule=diff is specified, the diff format is used. This format shows an inline diff of the changes in the submodule contents between the commit range. Defaults to diff.submodule or the short format if the config option is unset.

--color[=<when>]
Show colored diff. --color (i.e. without =<when>) is the same as --color=always. <when> can be one of always, never, or auto. It can be changed by the color.ui and color.diff configuration settings.

--no-color
Turn off colored diff. This can be used to override configuration settings. It is the same as --color=never.

--color-moved[=<mode>]
Moved lines of code are colored differently. It can be changed by the diff.colorMoved configuration setting. The <mode> defaults to no if the option is not given and to zebra if the option with no mode is given.

The mode must be one of:

no
Moved lines are not highlighted.

default
Is a synonym for zebra. This may change to a more sensible mode in the future.

plain
Any line that is added in one location and was removed in another location will be colored with color.diff.newMoved. Similarly color.diff.oldMoved will be used for removed lines that are added somewhere else in the diff. This mode picks up any moved line, but it is not very useful in a review to determine if a block of code was moved without permutation.

blocks
Blocks of moved text of at least 20 alphanumeric characters are detected greedily. The detected blocks are painted using either the color.diff.{old,new}Moved color. Adjacent blocks cannot be told apart.

zebra
Blocks of moved text are detected as in blocks mode. The blocks are painted using either the color.diff.{old,new}Moved color or color.diff.{old,new}MovedAlternative. The change between the two colors indicates that a new block was detected.

dimmed-zebra
Similar to zebra, but additional dimming of uninteresting parts of moved code is performed. The bordering lines of two adjacent blocks are considered interesting, the rest is uninteresting. dimmed_zebra is a deprecated synonym.

--no-color-moved
Turn off move detection. This can be used to override configuration settings. It is the same as --color-moved=no.

--color-moved-ws=<modes>
This configures how whitespace is ignored when performing the move detection for --color-moved. It can be set by the diff.colorMovedWS configuration setting.

These modes can be given as a comma separated list:

no
Do not ignore whitespace when performing move detection.

ignore-space-at-eol
Ignore changes in whitespace at EOL.

ignore-space-change
Ignore changes in amount of whitespace. This ignores whitespace at line end, and considers all other sequences of one or more whitespace characters to be equivalent.

ignore-all-space
Ignore whitespace when comparing lines. This ignores differences even if one line has whitespace where the other line has none.

allow-indentation-change
Initially ignore any whitespace in the move detection, then group the moved code blocks only into a block if the change in whitespace is the same per line. This is incompatible with the other modes.

--no-color-moved-ws
Do not ignore whitespace when performing move detection. This can be used to override configuration settings. It is the same as --color-moved-ws=no.

--word-diff[=<mode>]
Show a word diff, using the <mode> to delimit changed words. By default, words are delimited by whitespace; see --word-diff-regex below.

The <mode> defaults to plain, and must be one of:

color
Highlight changed words using only colors. Implies --color.

plain
Show words as [-removed-] and {+added+}. Makes no attempts to escape the delimiters if they appear in the input, so the output may be ambiguous.

porcelain
Use a special line-based format intended for script consumption. Added/removed/unchanged runs are printed in the usual unified diff format.

none
Disable word diff again.

Note that despite the name of the first mode, color is used to highlight the changed parts in all modes if enabled.

--word-diff-regex=<regex>
Use <regex> to decide what a word is, instead of considering runs of non-whitespace to be a word. Also implies --word-diff unless it was already enabled.

Every non-overlapping match of the <regex> is considered a word. Anything between these matches is considered whitespace and ignored(!) for the purposes of finding differences. You may want to append |[^[:space:]] to your regular expression to make sure that it matches all non-whitespace characters. A match that contains a newline is silently truncated(!) at the newline.

For example, --word-diff-regex=. will treat each character as a word and, correspondingly, show differences character by character.

The regex can also be set via a diff driver or configuration option, see gitattributes[5] or git-config[1]. Giving it explicitly overrides any diff driver or configuration setting. Diff drivers override configuration settings.

--color-words[=<regex>]
Equivalent to --word-diff=color plus (if a regex was specified) --word-diff-regex=<regex>.

--no-renames
Turn off rename detection, even when the configuration file gives the default to do so.

--[no-]rename-empty
Whether to use empty blobs as rename source.

--check
Warn if changes introduce conflict markers or whitespace errors. What are considered whitespace errors is controlled by core.whitespace configuration. By default, trailing whitespaces (including lines that consist solely of whitespaces) and a space character that is immediately followed by a tab character inside the initial indent of the line are considered whitespace errors. Exits with non-zero status if problems are found. Not compatible with --exit-code.

--ws-error-highlight=<kind>
Highlight whitespace errors in the context, old or new lines of the diff. Multiple values are separated by comma, none resets previous values, default reset the list to new and all is a shorthand for old,new,context. When this option is not given, and the configuration variable diff.wsErrorHighlight is not set, only whitespace errors in new lines are highlighted. The whitespace errors are colored with color.diff.whitespace.

--full-index
Instead of the first handful of characters, show the full pre- and post-image blob object names on the "index" line when generating patch format output.

--binary
In addition to --full-index, output a binary diff that can be applied with git-apply. Implies --patch.

--abbrev[=<n>]
Instead of showing the full 40-byte hexadecimal object name in diff-raw format output and diff-tree header lines, show only a partial prefix. This is independent of the --full-index option above, which controls the diff-patch output format. Non default number of digits can be specified with --abbrev=<n>.

-B[<n>][/<m>]
--break-rewrites[=[<n>][/<m>]]
Break complete rewrite changes into pairs of delete and create.

This serves two purposes:

It affects the way a change that amounts to a total rewrite of a file not as a series of deletion and insertion mixed together with a very few lines that happen to match textually as the context, but as a single deletion of everything old followed by a single insertion of everything new, and the number m controls this aspect of the -B option (defaults to 60%). -B/70% specifies that less than 30% of the original should remain in the result for Git to consider it a total rewrite (i.e. otherwise the resulting patch will be a series of deletion and insertion mixed together with context lines).

When used with -M, a totally-rewritten file is also considered as the source of a rename (usually -M only considers a file that disappeared as the source of a rename), and the number n controls this aspect of the -B option (defaults to 50%). -B20% specifies that a change with addition and deletion compared to 20% or more of the file’s size are eligible for being picked up as a possible source of a rename to another file.

-M[<n>]
--find-renames[=<n>]
Detect renames. If n is specified, it is a threshold on the similarity index (i.e. amount of addition/deletions compared to the file’s size). For example, -M90% means Git should consider a delete/add pair to be a rename if more than 90% of the file hasn’t changed. Without a % sign, the number is to be read as a fraction, with a decimal point before it. I.e., -M5 becomes 0.5, and is thus the same as -M50%. Similarly, -M05 is the same as -M5%. To limit detection to exact renames, use -M100%. The default similarity index is 50%.

-C[<n>]
--find-copies[=<n>]
Detect copies as well as renames. See also --find-copies-harder. If n is specified, it has the same meaning as for -M<n>.

--find-copies-harder
For performance reasons, by default, -C option finds copies only if the original file of the copy was modified in the same changeset. This flag makes the command inspect unmodified files as candidates for the source of copy. This is a very expensive operation for large projects, so use it with caution. Giving more than one -C option has the same effect.

-D
--irreversible-delete
Omit the preimage for deletes, i.e. print only the header but not the diff between the preimage and /dev/null. The resulting patch is not meant to be applied with patch or git apply; this is solely for people who want to just concentrate on reviewing the text after the change. In addition, the output obviously lacks enough information to apply such a patch in reverse, even manually, hence the name of the option.

When used together with -B, omit also the preimage in the deletion part of a delete/create pair.

-l<num>
The -M and -C options require O(n^2) processing time where n is the number of potential rename/copy targets. This option prevents rename/copy detection from running if the number of rename/copy targets exceeds the specified number.

--diff-filter=[(A|C|D|M|R|T|U|X|B)…​[*]]
Select only files that are Added (A), Copied (C), Deleted (D), Modified (M), Renamed (R), have their type (i.e. regular file, symlink, submodule, …​) changed (T), are Unmerged (U), are Unknown (X), or have had their pairing Broken (B). Any combination of the filter characters (including none) can be used. When * (All-or-none) is added to the combination, all paths are selected if there is any file that matches other criteria in the comparison; if there is no file that matches other criteria, nothing is selected.

Also, these upper-case letters can be downcased to exclude. E.g. --diff-filter=ad excludes added and deleted paths.

Note that not all diffs can feature all types. For instance, diffs from the index to the working tree can never have Added entries (because the set of paths included in the diff is limited by what is in the index). Similarly, copied and renamed entries cannot appear if detection for those types is disabled.

-S<string>
Look for differences that change the number of occurrences of the specified string (i.e. addition/deletion) in a file. Intended for the scripter’s use.

It is useful when you’re looking for an exact block of code (like a struct), and want to know the history of that block since it first came into being: use the feature iteratively to feed the interesting block in the preimage back into -S, and keep going until you get the very first version of the block.

Binary files are searched as well.

-G<regex>
Look for differences whose patch text contains added/removed lines that match <regex>.

To illustrate the difference between -S<regex> --pickaxe-regex and -G<regex>, consider a commit with the following diff in the same file:

+    return frotz(nitfol, two->ptr, 1, 0);
...
-    hit = frotz(nitfol, mf2.ptr, 1, 0);
While git log -G"frotz\(nitfol" will show this commit, git log -S"frotz\(nitfol" --pickaxe-regex will not (because the number of occurrences of that string did not change).

Unless --text is supplied patches of binary files without a textconv filter will be ignored.

See the pickaxe entry in gitdiffcore[7] for more information.

--find-object=<object-id>
Look for differences that change the number of occurrences of the specified object. Similar to -S, just the argument is different in that it doesn’t search for a specific string but for a specific object id.

The object can be a blob or a submodule commit. It implies the -t option in git-log to also find trees.

--pickaxe-all
When -S or -G finds a change, show all the changes in that changeset, not just the files that contain the change in <string>.

--pickaxe-regex
Treat the <string> given to -S as an extended POSIX regular expression to match.

-O<orderfile>
Control the order in which files appear in the output. This overrides the diff.orderFile configuration variable (see git-config[1]). To cancel diff.orderFile, use -O/dev/null.

The output order is determined by the order of glob patterns in <orderfile>. All files with pathnames that match the first pattern are output first, all files with pathnames that match the second pattern (but not the first) are output next, and so on. All files with pathnames that do not match any pattern are output last, as if there was an implicit match-all pattern at the end of the file. If multiple pathnames have the same rank (they match the same pattern but no earlier patterns), their output order relative to each other is the normal order.

<orderfile> is parsed as follows:

Blank lines are ignored, so they can be used as separators for readability.

Lines starting with a hash ("#") are ignored, so they can be used for comments. Add a backslash ("\") to the beginning of the pattern if it starts with a hash.

Each other line contains a single pattern.

Patterns have the same syntax and semantics as patterns used for fnmatch(3) without the FNM_PATHNAME flag, except a pathname also matches a pattern if removing any number of the final pathname components matches the pattern. For example, the pattern "foo*bar" matches "fooasdfbar" and "foo/bar/baz/asdf" but not "foobarx".

-R
Swap two inputs; that is, show differences from index or on-disk file to tree contents.

--relative[=<path>]
--no-relative
When run from a subdirectory of the project, it can be told to exclude changes outside the directory and show pathnames relative to it with this option. When you are not in a subdirectory (e.g. in a bare repository), you can name which subdirectory to make the output relative to by giving a <path> as an argument. --no-relative can be used to countermand both diff.relative config option and previous --relative.

-a
--text
Treat all files as text.

--ignore-cr-at-eol
Ignore carriage-return at the end of line when doing a comparison.

--ignore-space-at-eol
Ignore changes in whitespace at EOL.

-b
--ignore-space-change
Ignore changes in amount of whitespace. This ignores whitespace at line end, and considers all other sequences of one or more whitespace characters to be equivalent.

-w
--ignore-all-space
Ignore whitespace when comparing lines. This ignores differences even if one line has whitespace where the other line has none.

--ignore-blank-lines
Ignore changes whose lines are all blank.

--inter-hunk-context=<lines>
Show the context between diff hunks, up to the specified number of lines, thereby fusing hunks that are close to each other. Defaults to diff.interHunkContext or 0 if the config option is unset.

-W
--function-context
Show whole surrounding functions of changes.

--exit-code
Make the program exit with codes similar to diff(1). That is, it exits with 1 if there were differences and 0 means no differences.

--quiet
Disable all output of the program. Implies --exit-code.

--ext-diff
Allow an external diff helper to be executed. If you set an external diff driver with gitattributes[5], you need to use this option with git-log[1] and friends.

--no-ext-diff
Disallow external diff drivers.

--textconv
--no-textconv
Allow (or disallow) external text conversion filters to be run when comparing binary files. See gitattributes[5] for details. Because textconv filters are typically a one-way conversion, the resulting diff is suitable for human consumption, but cannot be applied. For this reason, textconv filters are enabled by default only for git-diff[1] and git-log[1], but not for git-format-patch[1] or diff plumbing commands.

--ignore-submodules[=<when>]
Ignore changes to submodules in the diff generation. <when> can be either "none", "untracked", "dirty" or "all", which is the default. Using "none" will consider the submodule modified when it either contains untracked or modified files or its HEAD differs from the commit recorded in the superproject and can be used to override any settings of the ignore option in git-config[1] or gitmodules[5]. When "untracked" is used submodules are not considered dirty when they only contain untracked content (but they are still scanned for modified content). Using "dirty" ignores all changes to the work tree of submodules, only changes to the commits stored in the superproject are shown (this was the behavior until 1.7.0). Using "all" hides all changes to submodules.

--src-prefix=<prefix>
Show the given source prefix instead of "a/".

--dst-prefix=<prefix>
Show the given destination prefix instead of "b/".

--no-prefix
Do not show any source or destination prefix.

--line-prefix=<prefix>
Prepend an additional prefix to every line of output.

--ita-invisible-in-index
By default entries added by "git add -N" appear as an existing empty file in "git diff" and a new file in "git diff --cached". This option makes the entry appear as a new file in "git diff" and non-existent in "git diff --cached". This option could be reverted with --ita-visible-in-index. Both options are experimental and could be removed in future.

For more detailed explanation on these common options, see also gitdiffcore[7].

-1 --base
-2 --ours
-3 --theirs
Compare the working tree with the "base" version (stage #1), "our branch" (stage #2) or "their branch" (stage #3). The index contains these stages only for unmerged entries i.e. while resolving conflicts. See git-read-tree[1] section "3-Way Merge" for detailed information.

-0
Omit diff output for unmerged entries and just show "Unmerged". Can be used only when comparing the working tree with the index.

<path>…​
The <paths> parameters, when given, are used to limit the diff to the named paths (you can give directory names and get diff for all files under them).`

const statusFlagsStr = `-s
--short
Give the output in the short-format.

-b
--branch
Show the branch and tracking info even in short-format.

--show-stash
Show the number of entries currently stashed away.

--porcelain[=<version>]
Give the output in an easy-to-parse format for scripts. This is similar to the short output, but will remain stable across Git versions and regardless of user configuration. See below for details.

The version parameter is used to specify the format version. This is optional and defaults to the original version v1 format.

--long
Give the output in the long-format. This is the default.

-v
--verbose
In addition to the names of files that have been changed, also show the textual changes that are staged to be committed (i.e., like the output of git diff --cached). If -v is specified twice, then also show the changes in the working tree that have not yet been staged (i.e., like the output of git diff).

-u[<mode>]
--untracked-files[=<mode>]
Show untracked files.

The mode parameter is used to specify the handling of untracked files. It is optional: it defaults to all, and if specified, it must be stuck to the option (e.g. -uno, but not -u no).

The possible options are:

no - Show no untracked files.

normal - Shows untracked files and directories.

all - Also shows individual files in untracked directories.

When -u option is not used, untracked files and directories are shown (i.e. the same as specifying normal), to help you avoid forgetting to add newly created files. Because it takes extra work to find untracked files in the filesystem, this mode may take some time in a large working tree. Consider enabling untracked cache and split index if supported (see git update-index --untracked-cache and git update-index --split-index), Otherwise you can use no to have git status return more quickly without showing untracked files.

The default can be changed using the status.showUntrackedFiles configuration variable documented in git-config[1].

--ignore-submodules[=<when>]
Ignore changes to submodules when looking for changes. <when> can be either "none", "untracked", "dirty" or "all", which is the default. Using "none" will consider the submodule modified when it either contains untracked or modified files or its HEAD differs from the commit recorded in the superproject and can be used to override any settings of the ignore option in git-config[1] or gitmodules[5]. When "untracked" is used submodules are not considered dirty when they only contain untracked content (but they are still scanned for modified content). Using "dirty" ignores all changes to the work tree of submodules, only changes to the commits stored in the superproject are shown (this was the behavior before 1.7.0). Using "all" hides all changes to submodules (and suppresses the output of submodule summaries when the config option status.submoduleSummary is set).

--ignored[=<mode>]
Show ignored files as well.

The mode parameter is used to specify the handling of ignored files. It is optional: it defaults to traditional.

The possible options are:

traditional - Shows ignored files and directories, unless --untracked-files=all is specified, in which case individual files in ignored directories are displayed.

no - Show no ignored files.

matching - Shows ignored files and directories matching an ignore pattern.

When matching mode is specified, paths that explicitly match an ignored pattern are shown. If a directory matches an ignore pattern, then it is shown, but not paths contained in the ignored directory. If a directory does not match an ignore pattern, but all contents are ignored, then the directory is not shown, but all contents are shown.

-z
Terminate entries with NUL, instead of LF. This implies the --porcelain=v1 output format if no other format is given.

--column[=<options>]
--no-column
Display untracked files in columns. See configuration variable column.status for option syntax.--column and --no-column without options are equivalent to always and never respectively.

--ahead-behind
--no-ahead-behind
Display or do not display detailed ahead/behind counts for the branch relative to its upstream branch. Defaults to true.

--renames
--no-renames
Turn on/off rename detection regardless of user configuration. See also git-diff[1] --no-renames.

--find-renames[=<n>]
Turn on rename detection, optionally setting the similarity threshold. See also git-diff[1] --find-renames.`

const commitFlagsStr = `-a
--all
Tell the command to automatically stage files that have been modified and deleted, but new files you have not told Git about are not affected.

-p
--patch
Use the interactive patch selection interface to chose which changes to commit. See git-add[1] for details.

-C <commit>
--reuse-message=<commit>
Take an existing commit object, and reuse the log message and the authorship information (including the timestamp) when creating the commit.

-c <commit>
--reedit-message=<commit>
Like -C, but with -c the editor is invoked, so that the user can further edit the commit message.

--fixup=<commit>
Construct a commit message for use with rebase --autosquash. The commit message will be the subject line from the specified commit with a prefix of "fixup! ". See git-rebase[1] for details.

--squash=<commit>
Construct a commit message for use with rebase --autosquash. The commit message subject line is taken from the specified commit with a prefix of "squash! ". Can be used with additional commit message options (-m/-c/-C/-F). See git-rebase[1] for details.

--reset-author
When used with -C/-c/--amend options, or when committing after a conflicting cherry-pick, declare that the authorship of the resulting commit now belongs to the committer. This also renews the author timestamp.

--short
When doing a dry-run, give the output in the short-format. See git-status[1] for details. Implies --dry-run.

--branch
Show the branch and tracking info even in short-format.

--porcelain
When doing a dry-run, give the output in a porcelain-ready format. See git-status[1] for details. Implies --dry-run.

--long
When doing a dry-run, give the output in the long-format. Implies --dry-run.

-z
--null
When showing short or porcelain status output, print the filename verbatim and terminate the entries with NUL, instead of LF. If no format is given, implies the --porcelain output format. Without the -z option, filenames with "unusual" characters are quoted as explained for the configuration variable core.quotePath (see git-config[1]).

-F <file>
--file=<file>
Take the commit message from the given file. Use - to read the message from the standard input.

--author=<author>
Override the commit author. Specify an explicit author using the standard A U Thor <author@example.com> format. Otherwise <author> is assumed to be a pattern and is used to search for an existing commit by that author (i.e. rev-list --all -i --author=<author>); the commit author is then copied from the first such commit found.

--date=<date>
Override the author date used in the commit.

-m <msg>
--message=<msg>
Use the given <msg> as the commit message. If multiple -m options are given, their values are concatenated as separate paragraphs.

The -m option is mutually exclusive with -c, -C, and -F.

-t <file>
--template=<file>
When editing the commit message, start the editor with the contents in the given file. The commit.template configuration variable is often used to give this option implicitly to the command. This mechanism can be used by projects that want to guide participants with some hints on what to write in the message in what order. If the user exits the editor without editing the message, the commit is aborted. This has no effect when a message is given by other means, e.g. with the -m or -F options.

-s
--signoff
Add Signed-off-by line by the committer at the end of the commit log message. The meaning of a signoff depends on the project, but it typically certifies that committer has the rights to submit this work under the same license and agrees to a Developer Certificate of Origin (see http://developercertificate.org/ for more information).

-n
--no-verify
This option bypasses the pre-commit and commit-msg hooks. See also githooks[5].

--allow-empty
Usually recording a commit that has the exact same tree as its sole parent commit is a mistake, and the command prevents you from making such a commit. This option bypasses the safety, and is primarily for use by foreign SCM interface scripts.

--allow-empty-message
Like --allow-empty this command is primarily for use by foreign SCM interface scripts. It allows you to create a commit with an empty commit message without using plumbing commands like git-commit-tree[1].

--cleanup=<mode>
This option determines how the supplied commit message should be cleaned up before committing. The <mode> can be strip, whitespace, verbatim, scissors or default.

strip
Strip leading and trailing empty lines, trailing whitespace, commentary and collapse consecutive empty lines.

whitespace
Same as strip except #commentary is not removed.

verbatim
Do not change the message at all.

scissors
Same as whitespace except that everything from (and including) the line found below is truncated, if the message is to be edited. "#" can be customized with core.commentChar.

# ------------------------ >8 ------------------------
default
Same as strip if the message is to be edited. Otherwise whitespace.

The default can be changed by the commit.cleanup configuration variable (see git-config[1]).

-e
--edit
The message taken from file with -F, command line with -m, and from commit object with -C are usually used as the commit log message unmodified. This option lets you further edit the message taken from these sources.

--no-edit
Use the selected commit message without launching an editor. For example, git commit --amend --no-edit amends a commit without changing its commit message.

--amend
Replace the tip of the current branch by creating a new commit. The recorded tree is prepared as usual (including the effect of the -i and -o options and explicit pathspec), and the message from the original commit is used as the starting point, instead of an empty message, when no other message is specified from the command line via options such as -m, -F, -c, etc. The new commit has the same parents and author as the current one (the --reset-author option can countermand this).

It is a rough equivalent for:

	$ git reset --soft HEAD^
	$ ... do something else to come up with the right tree ...
	$ git commit -c ORIG_HEAD
but can be used to amend a merge commit.

You should understand the implications of rewriting history if you amend a commit that has already been published. (See the "RECOVERING FROM UPSTREAM REBASE" section in git-rebase[1].)

--no-post-rewrite
Bypass the post-rewrite hook.

-i
--include
Before making a commit out of staged contents so far, stage the contents of paths given on the command line as well. This is usually not what you want unless you are concluding a conflicted merge.

-o
--only
Make a commit by taking the updated working tree contents of the paths specified on the command line, disregarding any contents that have been staged for other paths. This is the default mode of operation of git commit if any paths are given on the command line, in which case this option can be omitted. If this option is specified together with --amend, then no paths need to be specified, which can be used to amend the last commit without committing changes that have already been staged. If used together with --allow-empty paths are also not required, and an empty commit will be created.

--pathspec-from-file=<file>
Pathspec is passed in <file> instead of commandline args. If <file> is exactly - then standard input is used. Pathspec elements are separated by LF or CR/LF. Pathspec elements can be quoted as explained for the configuration variable core.quotePath (see git-config[1]). See also --pathspec-file-nul and global --literal-pathspecs.

--pathspec-file-nul
Only meaningful with --pathspec-from-file. Pathspec elements are separated with NUL character and all other characters are taken literally (including newlines and quotes).

-u[<mode>]
--untracked-files[=<mode>]
Show untracked files.

The mode parameter is optional (defaults to all), and is used to specify the handling of untracked files; when -u is not used, the default is normal, i.e. show untracked files and directories.

The possible options are:

no - Show no untracked files

normal - Shows untracked files and directories

all - Also shows individual files in untracked directories.

The default can be changed using the status.showUntrackedFiles configuration variable documented in git-config[1].

-v
--verbose
Show unified diff between the HEAD commit and what would be committed at the bottom of the commit message template to help the user describe the commit by reminding what changes the commit has. Note that this diff output doesn’t have its lines prefixed with #. This diff will not be a part of the commit message. See the commit.verbose configuration variable in git-config[1].

If specified twice, show in addition the unified diff between what would be committed and the worktree files, i.e. the unstaged changes to tracked files.

-q
--quiet
Suppress commit summary message.

--dry-run
Do not create a commit, but show a list of paths that are to be committed, paths with local changes that will be left uncommitted and paths that are untracked.

--status
Include the output of git-status[1] in the commit message template when using an editor to prepare the commit message. Defaults to on, but can be used to override configuration variable commit.status.

--no-status
Do not include the output of git-status[1] in the commit message template when using an editor to prepare the default commit message.

-S[<keyid>]
--gpg-sign[=<keyid>]
--no-gpg-sign
GPG-sign commits. The keyid argument is optional and defaults to the committer identity; if specified, it must be stuck to the option without a space. --no-gpg-sign is useful to countermand both commit.gpgSign configuration variable, and earlier --gpg-sign.

--
Do not interpret any more arguments as options.

<pathspec>…​
When pathspec is given on the command line, commit the contents of the files that match the pathspec without recording the changes already added to the index. The contents of these files are also staged for the next commit on top of what have been staged before.
`

const branchFlagsStr = `-d
--delete
Delete a branch. The branch must be fully merged in its upstream branch, or in HEAD if no upstream was set with --track or --set-upstream-to.

-D
Shortcut for --delete --force.

--create-reflog
Create the branch’s reflog. This activates recording of all changes made to the branch ref, enabling use of date based sha1 expressions such as "<branchname>@{yesterday}". Note that in non-bare repositories, reflogs are usually enabled by default by the core.logAllRefUpdates config option. The negated form --no-create-reflog only overrides an earlier --create-reflog, but currently does not negate the setting of core.logAllRefUpdates.

-f
--force
Reset <branchname> to <startpoint>, even if <branchname> exists already. Without -f, git branch refuses to change an existing branch. In combination with -d (or --delete), allow deleting the branch irrespective of its merged status. In combination with -m (or --move), allow renaming the branch even if the new branch name already exists, the same applies for -c (or --copy).

-m
--move
Move/rename a branch and the corresponding reflog.

-M
Shortcut for --move --force.

-c
--copy
Copy a branch and the corresponding reflog.

-C
Shortcut for --copy --force.

--color[=<when>]
Color branches to highlight current, local, and remote-tracking branches. The value must be always (the default), never, or auto.

--no-color
Turn off branch colors, even when the configuration file gives the default to color output. Same as --color=never.

-i
--ignore-case
Sorting and filtering branches are case insensitive.

--column[=<options>]
--no-column
Display branch listing in columns. See configuration variable column.branch for option syntax.--column and --no-column without options are equivalent to always and never respectively.

This option is only applicable in non-verbose mode.

-r
--remotes
List or delete (if used with -d) the remote-tracking branches. Combine with --list to match the optional pattern(s).

-a
--all
List both remote-tracking branches and local branches. Combine with --list to match optional pattern(s).

-l
--list
List branches. With optional <pattern>..., e.g. git branch --list 'maint-*', list only the branches that match the pattern(s).

--show-current
Print the name of the current branch. In detached HEAD state, nothing is printed.

-v
-vv
--verbose
When in list mode, show sha1 and commit subject line for each head, along with relationship to upstream branch (if any). If given twice, print the path of the linked worktree (if any) and the name of the upstream branch, as well (see also git remote show <remote>). Note that the current worktree’s HEAD will not have its path printed (it will always be your current directory).

-q
--quiet
Be more quiet when creating or deleting a branch, suppressing non-error messages.

--abbrev=<length>
Alter the sha1’s minimum display length in the output listing. The default value is 7 and can be overridden by the core.abbrev config option.

--no-abbrev
Display the full sha1s in the output listing rather than abbreviating them.

-t
--track
When creating a new branch, set up branch.<name>.remote and branch.<name>.merge configuration entries to mark the start-point branch as "upstream" from the new branch. This configuration will tell git to show the relationship between the two branches in git status and git branch -v. Furthermore, it directs git pull without arguments to pull from the upstream when the new branch is checked out.

This behavior is the default when the start point is a remote-tracking branch. Set the branch.autoSetupMerge configuration variable to false if you want git switch, git checkout and git branch to always behave as if --no-track were given. Set it to always if you want this behavior when the start-point is either a local or remote-tracking branch.

--no-track
Do not set up "upstream" configuration, even if the branch.autoSetupMerge configuration variable is true.

--set-upstream
As this option had confusing syntax, it is no longer supported. Please use --track or --set-upstream-to instead.

-u <upstream>
--set-upstream-to=<upstream>
Set up <branchname>'s tracking information so <upstream> is considered <branchname>'s upstream branch. If no <branchname> is specified, then it defaults to the current branch.

--unset-upstream
Remove the upstream information for <branchname>. If no branch is specified it defaults to the current branch.

--edit-description
Open an editor and edit the text to explain what the branch is for, to be used by various other commands (e.g. format-patch, request-pull, and merge (if enabled)). Multi-line explanations may be used.

--contains [<commit>]
Only list branches which contain the specified commit (HEAD if not specified). Implies --list.

--no-contains [<commit>]
Only list branches which don’t contain the specified commit (HEAD if not specified). Implies --list.

--merged [<commit>]
Only list branches whose tips are reachable from the specified commit (HEAD if not specified). Implies --list, incompatible with --no-merged.

--no-merged [<commit>]
Only list branches whose tips are not reachable from the specified commit (HEAD if not specified). Implies --list, incompatible with --merged.

<branchname>
The name of the branch to create or delete. The new branch name must pass all checks defined by git-check-ref-format[1]. Some of these checks may restrict the characters allowed in a branch name.

<start-point>
The new branch head will point to this commit. It may be given as a branch name, a commit-id, or a tag. If this option is omitted, the current HEAD will be used instead.

<oldbranch>
The name of an existing branch to rename.

<newbranch>
The new name for an existing branch. The same restrictions as for <branchname> apply.

--sort=<key>
Sort based on the key given. Prefix - to sort in descending order of the value. You may use the --sort=<key> option multiple times, in which case the last key becomes the primary key. The keys supported are the same as those in git for-each-ref. Sort order defaults to the value configured for the branch.sort variable if exists, or to sorting based on the full refname (including refs/... prefix). This lists detached HEAD (if present) first, then local branches and finally remote-tracking branches. See git-config[1].

--points-at <object>
Only list branches of the given object.

--format <format>
A string that interpolates %(fieldname) from a branch ref being shown and the object it points at. The format is the same as that of git-for-each-ref[1].`

const tagFlagsStr = `-a
--annotate
Make an unsigned, annotated tag object.

-s
--sign
Make a GPG-signed tag, using the default e-mail address’s key. The default behavior of tag GPG-signing is controlled by tag.gpgSign configuration variable if it exists, or disabled otherwise. See git-config[1].

--no-sign
Override tag.gpgSign configuration variable that is set to force each and every tag to be signed.

-u <keyid>
--local-user=<keyid>
Make a GPG-signed tag, using the given key.

-f
--force
Replace an existing tag with the given name (instead of failing).

-d
--delete
Delete existing tags with the given names.

-v
--verify
Verify the GPG signature of the given tag names.

-n<num>
<num> specifies how many lines from the annotation, if any, are printed when using -l. Implies --list.

The default is not to print any annotation lines. If no number is given to -n, only the first line is printed. If the tag is not annotated, the commit message is displayed instead.

-l
--list
List tags. With optional <pattern>..., e.g. git tag --list 'v-*', list only the tags that match the pattern(s).

Running "git tag" without arguments also lists all tags. The pattern is a shell wildcard (i.e., matched using fnmatch(3)). Multiple patterns may be given; if any of them matches, the tag is shown.

This option is implicitly supplied if any other list-like option such as --contains is provided. See the documentation for each of those options for details.

--sort=<key>
Sort based on the key given. Prefix - to sort in descending order of the value. You may use the --sort=<key> option multiple times, in which case the last key becomes the primary key. Also supports "version:refname" or "v:refname" (tag names are treated as versions). The "version:refname" sort order can also be affected by the "versionsort.suffix" configuration variable. The keys supported are the same as those in git for-each-ref. Sort order defaults to the value configured for the tag.sort variable if it exists, or lexicographic order otherwise. See git-config[1].

--color[=<when>]
Respect any colors specified in the --format option. The <when> field must be one of always, never, or auto (if <when> is absent, behave as if always was given).

-i
--ignore-case
Sorting and filtering tags are case insensitive.

--column[=<options>]
--no-column
Display tag listing in columns. See configuration variable column.tag for option syntax.--column and --no-column without options are equivalent to always and never respectively.

This option is only applicable when listing tags without annotation lines.

--contains [<commit>]
Only list tags which contain the specified commit (HEAD if not specified). Implies --list.

--no-contains [<commit>]
Only list tags which don’t contain the specified commit (HEAD if not specified). Implies --list.

--merged [<commit>]
Only list tags whose commits are reachable from the specified commit (HEAD if not specified), incompatible with --no-merged.

--no-merged [<commit>]
Only list tags whose commits are not reachable from the specified commit (HEAD if not specified), incompatible with --merged.

--points-at <object>
Only list tags of the given object (HEAD if not specified). Implies --list.

-m <msg>
--message=<msg>
Use the given tag message (instead of prompting). If multiple -m options are given, their values are concatenated as separate paragraphs. Implies -a if none of -a, -s, or -u <keyid> is given.

-F <file>
--file=<file>
Take the tag message from the given file. Use - to read the message from the standard input. Implies -a if none of -a, -s, or -u <keyid> is given.

-e
--edit
The message taken from file with -F and command line with -m are usually used as the tag message unmodified. This option lets you further edit the message taken from these sources.

--cleanup=<mode>
This option sets how the tag message is cleaned up. The <mode> can be one of verbatim, whitespace and strip. The strip mode is default. The verbatim mode does not change message at all, whitespace removes just leading/trailing whitespace lines and strip removes both whitespace and commentary.

--create-reflog
Create a reflog for the tag. To globally enable reflogs for tags, see core.logAllRefUpdates in git-config[1]. The negated form --no-create-reflog only overrides an earlier --create-reflog, but currently does not negate the setting of core.logAllRefUpdates.

--format=<format>
A string that interpolates %(fieldname) from a tag ref being shown and the object it points at. The format is the same as that of git-for-each-ref[1]. When unspecified, defaults to %(refname:strip=2).

<tagname>
The name of the tag to create, delete, or describe. The new tag name must pass all checks defined by git-check-ref-format[1]. Some of these checks may restrict the characters allowed in a tag name.

<commit>
<object>
The object that the new tag will refer to, usually a commit. Defaults to HEAD.`

const checkoutFlagsStr = `-q
--quiet
Quiet, suppress feedback messages.

--progress
--no-progress
Progress status is reported on the standard error stream by default when it is attached to a terminal, unless --quiet is specified. This flag enables progress reporting even if not attached to a terminal, regardless of --quiet.

-f
--force
When switching branches, proceed even if the index or the working tree differs from HEAD. This is used to throw away local changes.

When checking out paths from the index, do not fail upon unmerged entries; instead, unmerged entries are ignored.

--ours
--theirs
When checking out paths from the index, check out stage #2 (ours) or #3 (theirs) for unmerged paths.

Note that during git rebase and git pull --rebase, ours and theirs may appear swapped; --ours gives the version from the branch the changes are rebased onto, while --theirs gives the version from the branch that holds your work that is being rebased.

This is because rebase is used in a workflow that treats the history at the remote as the shared canonical one, and treats the work done on the branch you are rebasing as the third-party work to be integrated, and you are temporarily assuming the role of the keeper of the canonical history during the rebase. As the keeper of the canonical history, you need to view the history from the remote as ours (i.e. "our shared canonical history"), while what you did on your side branch as theirs (i.e. "one contributor’s work on top of it").

-b <new_branch>
Create a new branch named <new_branch> and start it at <start_point>; see git-branch[1] for details.

-B <new_branch>
Creates the branch <new_branch> and start it at <start_point>; if it already exists, then reset it to <start_point>. This is equivalent to running "git branch" with "-f"; see git-branch[1] for details.

-t
--track
When creating a new branch, set up "upstream" configuration. See "--track" in git-branch[1] for details.

If no -b option is given, the name of the new branch will be derived from the remote-tracking branch, by looking at the local part of the refspec configured for the corresponding remote, and then stripping the initial part up to the "*". This would tell us to use hack as the local branch when branching off of origin/hack (or remotes/origin/hack, or even refs/remotes/origin/hack). If the given name has no slash, or the above guessing results in an empty name, the guessing is aborted. You can explicitly give a name with -b in such a case.

--no-track
Do not set up "upstream" configuration, even if the branch.autoSetupMerge configuration variable is true.

--guess
--no-guess
If <branch> is not found but there does exist a tracking branch in exactly one remote (call it <remote>) with a matching name, treat as equivalent to.

-l
Create the new branch’s reflog; see git-branch[1] for details.

--detach
Rather than checking out a branch to work on it, check out a commit for inspection and discardable experiments. This is the default behavior of git checkout <commit> when <commit> is not a branch name. See the "DETACHED HEAD" section below for details.

--orphan <new_branch>
Create a new orphan branch, named <new_branch>, started from <start_point> and switch to it. The first commit made on this new branch will have no parents and it will be the root of a new history totally disconnected from all the other branches and commits.

The index and the working tree are adjusted as if you had previously run git checkout <start_point>. This allows you to start a new history that records a set of paths similar to <start_point> by easily running git commit -a to make the root commit.

This can be useful when you want to publish the tree from a commit without exposing its full history. You might want to do this to publish an open source branch of a project whose current tree is "clean", but whose full history contains proprietary or otherwise encumbered bits of code.

If you want to start a disconnected history that records a set of paths that is totally different from the one of <start_point>, then you should clear the index and the working tree right after creating the orphan branch by running git rm -rf . from the top level of the working tree. Afterwards you will be ready to prepare your new files, repopulating the working tree, by copying them from elsewhere, extracting a tarball, etc.

--ignore-skip-worktree-bits
In sparse checkout mode, git checkout -- <paths> would update only entries matched by <paths> and sparse patterns in $GIT_DIR/info/sparse-checkout. This option ignores the sparse patterns and adds back any files in <paths>.

-m
--merge
When switching branches, if you have local modifications to one or more files that are different between the current branch and the branch to which you are switching, the command refuses to switch branches in order to preserve your modifications in context. However, with this option, a three-way merge between the current branch, your working tree contents, and the new branch is done, and you will be on the new branch.

When a merge conflict happens, the index entries for conflicting paths are left unmerged, and you need to resolve the conflicts and mark the resolved paths with git add (or git rm if the merge should result in deletion of the path).

When checking out paths from the index, this option lets you recreate the conflicted merge in the specified paths.

When switching branches with --merge, staged changes may be lost.

--conflict=<style>
The same as --merge option above, but changes the way the conflicting hunks are presented, overriding the merge.conflictStyle configuration variable. Possible values are "merge" (default) and "diff3" (in addition to what is shown by "merge" style, shows the original contents).

-p
--patch
Interactively select hunks in the difference between the <tree-ish> (or the index, if unspecified) and the working tree. The chosen hunks are then applied in reverse to the working tree (and if a <tree-ish> was specified, the index).

This means that you can use git checkout -p to selectively discard edits from your current working tree. See the “Interactive Mode” section of git-add[1] to learn how to operate the --patch mode.

Note that this option uses the no overlay mode by default (see also --overlay), and currently doesn’t support overlay mode.

--ignore-other-worktrees
git checkout refuses when the wanted ref is already checked out by another worktree. This option makes it check the ref out anyway. In other words, the ref can be held by more than one worktree.

--overwrite-ignore
--no-overwrite-ignore
Silently overwrite ignored files when switching branches. This is the default behavior. Use --no-overwrite-ignore to abort the operation when the new branch contains ignored files.

--recurse-submodules
--no-recurse-submodules
Using --recurse-submodules will update the content of all active submodules according to the commit recorded in the superproject. If local modifications in a submodule would be overwritten the checkout will fail unless -f is used. If nothing (or --no-recurse-submodules) is used, submodules working trees will not be updated. Just like git-submodule[1], this will detach HEAD of the submodule.

--overlay
--no-overlay
In the default overlay mode, git checkout never removes files from the index or the working tree. When specifying --no-overlay, files that appear in the index and working tree, but not in <tree-ish> are removed, to make them match <tree-ish> exactly.

--pathspec-from-file=<file>
Pathspec is passed in <file> instead of commandline args. If <file> is exactly - then standard input is used. Pathspec elements are separated by LF or CR/LF. Pathspec elements can be quoted as explained for the configuration variable core.quotePath (see git-config[1]). See also --pathspec-file-nul and global --literal-pathspecs.

--pathspec-file-nul
Only meaningful with --pathspec-from-file. Pathspec elements are separated with NUL character and all other characters are taken literally (including newlines and quotes).

<branch>
Branch to checkout; if it refers to a branch (i.e., a name that, when prepended with "refs/heads/", is a valid ref), then that branch is checked out. Otherwise, if it refers to a valid commit, your HEAD becomes "detached" and you are no longer on any branch (see below for details).

You can use the @{-N} syntax to refer to the N-th last branch/commit checked out using "git checkout" operation. You may also specify - which is synonymous to @{-1}.

As a special case, you may use A...B as a shortcut for the merge base of A and B if there is exactly one merge base. You can leave out at most one of A and B, in which case it defaults to HEAD.

<new_branch>
Name for the new branch.

<start_point>
The name of a commit at which to start the new branch; see git-branch[1] for details. Defaults to HEAD.

As a special case, you may use "A...B" as a shortcut for the merge base of A and B if there is exactly one merge base. You can leave out at most one of A and B, in which case it defaults to HEAD.

<tree-ish>
Tree to checkout from (when paths are given). If not specified, the index will be used.

--
Do not interpret any more arguments as options.

<pathspec>…​
Limits the paths affected by the operation.

For more details, see the pathspec entry in gitglossary[7].`

const mergeFlagsStr = `--commit
--no-commit
Perform the merge and commit the result. This option can be used to override --no-commit.

With --no-commit perform the merge and stop just before creating a merge commit, to give the user a chance to inspect and further tweak the merge result before committing.

Note that fast-forward updates do not create a merge commit and therefore there is no way to stop those merges with --no-commit. Thus, if you want to ensure your branch is not changed or updated by the merge command, use --no-ff with --no-commit.

--edit
-e
--no-edit
Invoke an editor before committing successful mechanical merge to further edit the auto-generated merge message, so that the user can explain and justify the merge. The --no-edit option can be used to accept the auto-generated message (this is generally discouraged). The --edit (or -e) option is still useful if you are giving a draft message with the -m option from the command line and want to edit it in the editor.

Older scripts may depend on the historical behaviour of not allowing the user to edit the merge log message. They will see an editor opened when they run git merge. To make it easier to adjust such scripts to the updated behaviour, the environment variable GIT_MERGE_AUTOEDIT can be set to no at the beginning of them.

--cleanup=<mode>
This option determines how the merge message will be cleaned up before committing. See git-commit[1] for more details. In addition, if the <mode> is given a value of scissors, scissors will be appended to MERGE_MSG before being passed on to the commit machinery in the case of a merge conflict.

--ff
--no-ff
--ff-only
Specifies how a merge is handled when the merged-in history is already a descendant of the current history. --ff is the default unless merging an annotated (and possibly signed) tag that is not stored in its natural place in the refs/tags/ hierarchy, in which case --no-ff is assumed.

With --ff, when possible resolve the merge as a fast-forward (only update the branch pointer to match the merged branch; do not create a merge commit). When not possible (when the merged-in history is not a descendant of the current history), create a merge commit.

With --no-ff, create a merge commit in all cases, even when the merge could instead be resolved as a fast-forward.

With --ff-only, resolve the merge as a fast-forward when possible. When not possible, refuse to merge and exit with a non-zero status.

-S[<keyid>]
--gpg-sign[=<keyid>]
--no-gpg-sign
GPG-sign the resulting merge commit. The keyid argument is optional and defaults to the committer identity; if specified, it must be stuck to the option without a space. --no-gpg-sign is useful to countermand both commit.gpgSign configuration variable, and earlier --gpg-sign.

--log[=<n>]
--no-log
In addition to branch names, populate the log message with one-line descriptions from at most <n> actual commits that are being merged. See also git-fmt-merge-msg[1].

With --no-log do not list one-line descriptions from the actual commits being merged.

--signoff
--no-signoff
Add Signed-off-by line by the committer at the end of the commit log message. The meaning of a signoff depends on the project, but it typically certifies that committer has the rights to submit this work under the same license and agrees to a Developer Certificate of Origin (see http://developercertificate.org/ for more information).

With --no-signoff do not add a Signed-off-by line.

--stat
-n
--no-stat
Show a diffstat at the end of the merge. The diffstat is also controlled by the configuration option merge.stat.

With -n or --no-stat do not show a diffstat at the end of the merge.

--squash
--no-squash
Produce the working tree and index state as if a real merge happened (except for the merge information), but do not actually make a commit, move the HEAD, or record $GIT_DIR/MERGE_HEAD (to cause the next git commit command to create a merge commit). This allows you to create a single commit on top of the current branch whose effect is the same as merging another branch (or more in case of an octopus).

With --no-squash perform the merge and commit the result. This option can be used to override --squash.

With --squash, --commit is not allowed, and will fail.

--no-verify
This option bypasses the pre-merge and commit-msg hooks. See also githooks[5].

-s <strategy>
--strategy=<strategy>
Use the given merge strategy; can be supplied more than once to specify them in the order they should be tried. If there is no -s option, a built-in list of strategies is used instead (git merge-recursive when merging a single head, git merge-octopus otherwise).

-X <option>
--strategy-option=<option>
Pass merge strategy specific option through to the merge strategy.

--verify-signatures
--no-verify-signatures
Verify that the tip commit of the side branch being merged is signed with a valid key, i.e. a key that has a valid uid: in the default trust model, this means the signing key has been signed by a trusted key. If the tip commit of the side branch is not signed with a valid key, the merge is aborted.

--summary
--no-summary
Synonyms to --stat and --no-stat; these are deprecated and will be removed in the future.

-q
--quiet
Operate quietly. Implies --no-progress.

-v
--verbose
Be verbose.

--progress
--no-progress
Turn progress on/off explicitly. If neither is specified, progress is shown if standard error is connected to a terminal. Note that not all merge strategies may support progress reporting.

--autostash
--no-autostash
Automatically create a temporary stash entry before the operation begins, and apply it after the operation ends. This means that you can run the operation on a dirty worktree. However, use with care: the final stash application after a successful merge might result in non-trivial conflicts.

--allow-unrelated-histories
By default, git merge command refuses to merge histories that do not share a common ancestor. This option can be used to override this safety when merging histories of two projects that started their lives independently. As that is a very rare occasion, no configuration variable to enable this by default exists and will not be added.

-m <msg>
Set the commit message to be used for the merge commit (in case one is created).

If --log is specified, a shortlog of the commits being merged will be appended to the specified message.

The git fmt-merge-msg command can be used to give a good default for automated git merge invocations. The automated message can include the branch description.

-F <file>
--file=<file>
Read the commit message to be used for the merge commit (in case one is created).

If --log is specified, a shortlog of the commits being merged will be appended to the specified message.

--rerere-autoupdate
--no-rerere-autoupdate
Allow the rerere mechanism to update the index with the result of auto-conflict resolution if possible.

--overwrite-ignore
--no-overwrite-ignore
Silently overwrite ignored files from the merge result. This is the default behavior. Use --no-overwrite-ignore to abort.

--abort
Abort the current conflict resolution process, and try to reconstruct the pre-merge state. If an autostash entry is present, apply it to the worktree.

If there were uncommitted worktree changes present when the merge started, git merge --abort will in some cases be unable to reconstruct these changes. It is therefore recommended to always commit or stash your changes before running git merge.

git merge --abort is equivalent to git reset --merge when MERGE_HEAD is present unless MERGE_AUTOSTASH is also present in which case git merge --abort applies the stash entry to the worktree whereas git reset --merge will save the stashed changes in the stash list.

--quit
Forget about the current merge in progress. Leave the index and the working tree as-is. If MERGE_AUTOSTASH is present, the stash entry will be saved to the stash list.

--continue
After a git merge stops due to conflicts you can conclude the merge by running git merge --continue (see "HOW TO RESOLVE CONFLICTS" section below).

<commit>…​
Commits, usually other branch heads, to merge into our branch. Specifying more than one commit will create a merge with more than two parents (affectionately called an Octopus merge).

If no commit is given from the command line, merge the remote-tracking branches that the current branch is configured to use as its upstream. See also the configuration section of this manual page.

When FETCH_HEAD (and no other commit) is specified, the branches recorded in the .git/FETCH_HEAD file by the previous invocation of git fetch for merging are merged to the current branch.`

const pullFlagsStr = `-q
--quiet
This is passed to both underlying git-fetch to squelch reporting of during transfer, and underlying git-merge to squelch output during merging.

-v
--verbose
Pass --verbose to git-fetch and git-merge.

--[no-]recurse-submodules[=yes|on-demand|no]
This option controls if new commits of populated submodules should be fetched, and if the working trees of active submodules should be updated, too (see git-fetch[1], git-config[1] and gitmodules[5]).

If the checkout is done via rebase, local submodule commits are rebased as well.

If the update is done via merge, the submodule conflicts are resolved and checked out.

Options related to merging
--commit
--no-commit
Perform the merge and commit the result. This option can be used to override --no-commit.

With --no-commit perform the merge and stop just before creating a merge commit, to give the user a chance to inspect and further tweak the merge result before committing.

Note that fast-forward updates do not create a merge commit and therefore there is no way to stop those merges with --no-commit. Thus, if you want to ensure your branch is not changed or updated by the merge command, use --no-ff with --no-commit.

--edit
-e
--no-edit
Invoke an editor before committing successful mechanical merge to further edit the auto-generated merge message, so that the user can explain and justify the merge. The --no-edit option can be used to accept the auto-generated message (this is generally discouraged).

Older scripts may depend on the historical behaviour of not allowing the user to edit the merge log message. They will see an editor opened when they run git merge. To make it easier to adjust such scripts to the updated behaviour, the environment variable GIT_MERGE_AUTOEDIT can be set to no at the beginning of them.

--cleanup=<mode>
This option determines how the merge message will be cleaned up before committing. See git-commit[1] for more details. In addition, if the <mode> is given a value of scissors, scissors will be appended to MERGE_MSG before being passed on to the commit machinery in the case of a merge conflict.

--ff
--no-ff
--ff-only
Specifies how a merge is handled when the merged-in history is already a descendant of the current history. --ff is the default unless merging an annotated (and possibly signed) tag that is not stored in its natural place in the refs/tags/ hierarchy, in which case --no-ff is assumed.

With --ff, when possible resolve the merge as a fast-forward (only update the branch pointer to match the merged branch; do not create a merge commit). When not possible (when the merged-in history is not a descendant of the current history), create a merge commit.

With --no-ff, create a merge commit in all cases, even when the merge could instead be resolved as a fast-forward.

With --ff-only, resolve the merge as a fast-forward when possible. When not possible, refuse to merge and exit with a non-zero status.

-S[<keyid>]
--gpg-sign[=<keyid>]
--no-gpg-sign
GPG-sign the resulting merge commit. The keyid argument is optional and defaults to the committer identity; if specified, it must be stuck to the option without a space. --no-gpg-sign is useful to countermand both commit.gpgSign configuration variable, and earlier --gpg-sign.

--log[=<n>]
--no-log
In addition to branch names, populate the log message with one-line descriptions from at most <n> actual commits that are being merged. See also git-fmt-merge-msg[1].

With --no-log do not list one-line descriptions from the actual commits being merged.

--signoff
--no-signoff
Add Signed-off-by line by the committer at the end of the commit log message. The meaning of a signoff depends on the project, but it typically certifies that committer has the rights to submit this work under the same license and agrees to a Developer Certificate of Origin (see http://developercertificate.org/ for more information).

With --no-signoff do not add a Signed-off-by line.

--stat
-n
--no-stat
Show a diffstat at the end of the merge. The diffstat is also controlled by the configuration option merge.stat.

With -n or --no-stat do not show a diffstat at the end of the merge.

--squash
--no-squash
Produce the working tree and index state as if a real merge happened (except for the merge information), but do not actually make a commit, move the HEAD, or record $GIT_DIR/MERGE_HEAD (to cause the next git commit command to create a merge commit). This allows you to create a single commit on top of the current branch whose effect is the same as merging another branch (or more in case of an octopus).

With --no-squash perform the merge and commit the result. This option can be used to override --squash.

With --squash, --commit is not allowed, and will fail.

--no-verify
This option bypasses the pre-merge and commit-msg hooks. See also githooks[5].

-s <strategy>
--strategy=<strategy>
Use the given merge strategy; can be supplied more than once to specify them in the order they should be tried. If there is no -s option, a built-in list of strategies is used instead (git merge-recursive when merging a single head, git merge-octopus otherwise).

-X <option>
--strategy-option=<option>
Pass merge strategy specific option through to the merge strategy.

--verify-signatures
--no-verify-signatures
Verify that the tip commit of the side branch being merged is signed with a valid key, i.e. a key that has a valid uid: in the default trust model, this means the signing key has been signed by a trusted key. If the tip commit of the side branch is not signed with a valid key, the merge is aborted.

--summary
--no-summary
Synonyms to --stat and --no-stat; these are deprecated and will be removed in the future.

--autostash
--no-autostash
Automatically create a temporary stash entry before the operation begins, and apply it after the operation ends. This means that you can run the operation on a dirty worktree. However, use with care: the final stash application after a successful merge might result in non-trivial conflicts.

--allow-unrelated-histories
By default, git merge command refuses to merge histories that do not share a common ancestor. This option can be used to override this safety when merging histories of two projects that started their lives independently. As that is a very rare occasion, no configuration variable to enable this by default exists and will not be added.

-r
--rebase[=false|true|merges|preserve|interactive]
When true, rebase the current branch on top of the upstream branch after fetching. If there is a remote-tracking branch corresponding to the upstream branch and the upstream branch was rebased since last fetched, the rebase uses that information to avoid rebasing non-local changes.

When set to merges, rebase using git rebase --rebase-merges so that the local merge commits are included in the rebase (see git-rebase[1] for details).

When set to preserve (deprecated in favor of merges), rebase with the --preserve-merges option passed to git rebase so that locally created merge commits will not be flattened.

When false, merge the current branch into the upstream branch.

When interactive, enable the interactive mode of rebase.

See pull.rebase, branch.<name>.rebase and branch.autoSetupRebase in git-config[1] if you want to make git pull always use --rebase instead of merging.

Note
This is a potentially dangerous mode of operation. It rewrites history, which does not bode well when you published that history already. Do not use this option unless you have read git-rebase[1] carefully.
--no-rebase
Override earlier --rebase.

Options related to fetching
--all
Fetch all remotes.

-a
--append
Append ref names and object names of fetched refs to the existing contents of .git/FETCH_HEAD. Without this option old data in .git/FETCH_HEAD will be overwritten.

--depth=<depth>
Limit fetching to the specified number of commits from the tip of each remote branch history. If fetching to a shallow repository created by git clone with --depth=<depth> option (see git-clone[1]), deepen or shorten the history to the specified number of commits. Tags for the deepened commits are not fetched.

--deepen=<depth>
Similar to --depth, except it specifies the number of commits from the current shallow boundary instead of from the tip of each remote branch history.

--shallow-since=<date>
Deepen or shorten the history of a shallow repository to include all reachable commits after <date>.

--shallow-exclude=<revision>
Deepen or shorten the history of a shallow repository to exclude commits reachable from a specified remote branch or tag. This option can be specified multiple times.

--unshallow
If the source repository is complete, convert a shallow repository to a complete one, removing all the limitations imposed by shallow repositories.

If the source repository is shallow, fetch as much as possible so that the current repository has the same history as the source repository.

--update-shallow
By default when fetching from a shallow repository, git fetch refuses refs that require updating .git/shallow. This option updates .git/shallow and accept such refs.

--negotiation-tip=<commit|glob>
By default, Git will report, to the server, commits reachable from all local refs to find common commits in an attempt to reduce the size of the to-be-received packfile. If specified, Git will only report commits reachable from the given tips. This is useful to speed up fetches when the user knows which local ref is likely to have commits in common with the upstream ref being fetched.

This option may be specified more than once; if so, Git will report commits reachable from any of the given commits.

The argument to this option may be a glob on ref names, a ref, or the (possibly abbreviated) SHA-1 of a commit. Specifying a glob is equivalent to specifying this option multiple times, one for each matching ref name.

See also the fetch.negotiationAlgorithm configuration variable documented in git-config[1].

--dry-run
Show what would be done, without making any changes.

-f
--force
When git fetch is used with <src>:<dst> refspec it may refuse to update the local branch as discussed in the <refspec> part of the git-fetch[1] documentation. This option overrides that check.

-k
--keep
Keep downloaded pack.

-p
--prune
Before fetching, remove any remote-tracking references that no longer exist on the remote. Tags are not subject to pruning if they are fetched only because of the default tag auto-following or due to a --tags option. However, if tags are fetched due to an explicit refspec (either on the command line or in the remote configuration, for example if the remote was cloned with the --mirror option), then they are also subject to pruning. Supplying --prune-tags is a shorthand for providing the tag refspec.

--no-tags
By default, tags that point at objects that are downloaded from the remote repository are fetched and stored locally. This option disables this automatic tag following. The default behavior for a remote may be specified with the remote.<name>.tagOpt setting. See git-config[1].

--refmap=<refspec>
When fetching refs listed on the command line, use the specified refspec (can be given more than once) to map the refs to remote-tracking branches, instead of the values of remote.*.fetch configuration variables for the remote repository. Providing an empty <refspec> to the --refmap option causes Git to ignore the configured refspecs and rely entirely on the refspecs supplied as command-line arguments. See section on "Configured Remote-tracking Branches" for details.

-t
--tags
Fetch all tags from the remote (i.e., fetch remote tags refs/tags/* into local tags with the same name), in addition to whatever else would otherwise be fetched. Using this option alone does not subject tags to pruning, even if --prune is used (though tags may be pruned anyway if they are also the destination of an explicit refspec; see --prune).

-j
--jobs=<n>
Number of parallel children to be used for all forms of fetching.

If the --multiple option was specified, the different remotes will be fetched in parallel. If multiple submodules are fetched, they will be fetched in parallel. To control them independently, use the config settings fetch.parallel and submodule.fetchJobs (see git-config[1]).

Typically, parallel recursive and multi-remote fetches will be faster. By default fetches are performed sequentially, not in parallel.

--set-upstream
If the remote is fetched successfully, pull and add upstream (tracking) reference, used by argument-less git-pull[1] and other commands. For more information, see branch.<name>.merge and branch.<name>.remote in git-config[1].

--upload-pack <upload-pack>
When given, and the repository to fetch from is handled by git fetch-pack, --exec=<upload-pack> is passed to the command to specify non-default path for the command run on the other end.

--progress
Progress status is reported on the standard error stream by default when it is attached to a terminal, unless -q is specified. This flag forces progress status even if the standard error stream is not directed to a terminal.

-o <option>
--server-option=<option>
Transmit the given string to the server when communicating using protocol version 2. The given string must not contain a NUL or LF character. The server’s handling of server options, including unknown ones, is server-specific. When multiple --server-option=<option> are given, they are all sent to the other side in the order listed on the command line.

--show-forced-updates
By default, git checks if a branch is force-updated during fetch. This can be disabled through fetch.showForcedUpdates, but the --show-forced-updates option guarantees this check occurs. See git-config[1].

--no-show-forced-updates
By default, git checks if a branch is force-updated during fetch. Pass --no-show-forced-updates or set fetch.showForcedUpdates to false to skip this check for performance reasons. If used during git-pull the --ff-only option will still check for forced updates before attempting a fast-forward update. See git-config[1].

-4
--ipv4
Use IPv4 addresses only, ignoring IPv6 addresses.

-6
--ipv6
Use IPv6 addresses only, ignoring IPv4 addresses.`

const pushFlagsStr = `
--all
Push all branches (i.e. refs under refs/heads/); cannot be used with other <refspec>.

--prune
Remove remote branches that don’t have a local counterpart. For example a remote branch tmp will be removed if a local branch with the same name doesn’t exist any more. This also respects refspecs, e.g. git push --prune remote refs/heads/*:refs/tmp/* would make sure that remote refs/tmp/foo will be removed if refs/heads/foo doesn’t exist.

--mirror
Instead of naming each ref to push, specifies that all refs under refs/ (which includes but is not limited to refs/heads/, refs/remotes/, and refs/tags/) be mirrored to the remote repository. Newly created local refs will be pushed to the remote end, locally updated refs will be force updated on the remote end, and deleted refs will be removed from the remote end. This is the default if the configuration option remote.<remote>.mirror is set.

-n
--dry-run
Do everything except actually send the updates.

--porcelain
Produce machine-readable output. The output status line for each ref will be tab-separated and sent to stdout instead of stderr. The full symbolic names of the refs will be given.

-d
--delete
All listed refs are deleted from the remote repository. This is the same as prefixing all refs with a colon.

--tags
All refs under refs/tags are pushed, in addition to refspecs explicitly listed on the command line.

--follow-tags
Push all the refs that would be pushed without this option, and also push annotated tags in refs/tags that are missing from the remote but are pointing at commit-ish that are reachable from the refs being pushed. This can also be specified with configuration variable push.followTags. For more information, see push.followTags in git-config[1].

--[no-]signed
--signed=(true|false|if-asked)
GPG-sign the push request to update refs on the receiving side, to allow it to be checked by the hooks and/or be logged. If false or --no-signed, no signing will be attempted. If true or --signed, the push will fail if the server does not support signed pushes. If set to if-asked, sign if and only if the server supports signed pushes. The push will also fail if the actual call to gpg --sign fails. See git-receive-pack[1] for the details on the receiving end.

--[no-]atomic
Use an atomic transaction on the remote side if available. Either all refs are updated, or on error, no refs are updated. If the server does not support atomic pushes the push will fail.

-o <option>
--push-option=<option>
Transmit the given string to the server, which passes them to the pre-receive as well as the post-receive hook. The given string must not contain a NUL or LF character. When multiple --push-option=<option> are given, they are all sent to the other side in the order listed on the command line. When no --push-option=<option> is given from the command line, the values of configuration variable push.pushOption are used instead.

--receive-pack=<git-receive-pack>
--exec=<git-receive-pack>
Path to the git-receive-pack program on the remote end. Sometimes useful when pushing to a remote repository over ssh, and you do not have the program in a directory on the default $PATH.

--[no-]force-with-lease
--force-with-lease
Usually, "git push" refuses to update a remote ref that is not an ancestor of the local ref used to overwrite it.

This option overrides this restriction if the current value of the remote ref is the expected value. "git push" fails otherwise.

Imagine that you have to rebase what you have already published. You will have to bypass the "must fast-forward" rule in order to replace the history you originally published with the rebased history. If somebody else built on top of your original history while you are rebasing, the tip of the branch at the remote may advance with her commit, and blindly pushing with --force will lose her work.

This option allows you to say that you expect the history you are updating is what you rebased and want to replace. If the remote ref still points at the commit you specified, you can be sure that no other people did anything to the ref. It is like taking a "lease" on the ref without explicitly locking it, and the remote ref is updated only if the "lease" is still valid.

--force-with-lease alone, without specifying the details, will protect all remote refs that are going to be updated by requiring their current value to be the same as the remote-tracking branch we have for them.

--force-with-lease=<refname>, without specifying the expected value, will protect the named ref (alone), if it is going to be updated, by requiring its current value to be the same as the remote-tracking branch we have for it.

--force-with-lease=<refname>:<expect> will protect the named ref (alone), if it is going to be updated, by requiring its current value to be the same as the specified value <expect> (which is allowed to be different from the remote-tracking branch we have for the refname, or we do not even have to have such a remote-tracking branch when this form is used). If <expect> is the empty string, then the named ref must not already exist.

Note that all forms other than --force-with-lease=<refname>:<expect> that specifies the expected current value of the ref explicitly are still experimental and their semantics may change as we gain experience with this feature.

"--no-force-with-lease" will cancel all the previous --force-with-lease on the command line.

A general note on safety: supplying this option without an expected value, i.e. as --force-with-lease or --force-with-lease=<refname> interacts very badly with anything that implicitly runs git fetch on the remote to be pushed to in the background, e.g. git fetch origin on your repository in a cronjob.

The protection it offers over --force is ensuring that subsequent changes your work wasn’t based on aren’t clobbered, but this is trivially defeated if some background process is updating refs in the background. We don’t have anything except the remote tracking info to go by as a heuristic for refs you’re expected to have seen & are willing to clobber.

If your editor or some other system is running git fetch in the background for you a way to mitigate this is to simply set up another remote:

git remote add origin-push $(git config remote.origin.url)
git fetch origin-push
Now when the background process runs git fetch origin the references on origin-push won’t be updated, and thus commands like:

git push --force-with-lease origin-push
Will fail unless you manually run git fetch origin-push. This method is of course entirely defeated by something that runs git fetch --all, in that case you’d need to either disable it or do something more tedious like:

git fetch              # update 'master' from remote
git tag base master    # mark our base point
git rebase -i master   # rewrite some commits
git push --force-with-lease=master:base master:master
I.e. create a base tag for versions of the upstream code that you’ve seen and are willing to overwrite, then rewrite history, and finally force push changes to master if the remote version is still at base, regardless of what your local remotes/origin/master has been updated to in the background.

-f
--force
Usually, the command refuses to update a remote ref that is not an ancestor of the local ref used to overwrite it. Also, when --force-with-lease option is used, the command refuses to update a remote ref whose current value does not match what is expected.

This flag disables these checks, and can cause the remote repository to lose commits; use it with care.

Note that --force applies to all the refs that are pushed, hence using it with push.default set to matching or with multiple push destinations configured with remote.*.push may overwrite refs other than the current branch (including local refs that are strictly behind their remote counterpart). To force a push to only one branch, use a + in front of the refspec to push (e.g git push origin +master to force a push to the master branch). See the <refspec>... section above for details.

--repo=<repository>
This option is equivalent to the <repository> argument. If both are specified, the command-line argument takes precedence.

-u
--set-upstream
For every branch that is up to date or successfully pushed, add upstream (tracking) reference, used by argument-less git-pull[1] and other commands. For more information, see branch.<name>.merge in git-config[1].

--[no-]thin
These options are passed to git-send-pack[1]. A thin transfer significantly reduces the amount of sent data when the sender and receiver share many of the same objects in common. The default is --thin.

-q
--quiet
Suppress all output, including the listing of updated refs, unless an error occurs. Progress is not reported to the standard error stream.

-v
--verbose
Run verbosely.

--progress
Progress status is reported on the standard error stream by default when it is attached to a terminal, unless -q is specified. This flag forces progress status even if the standard error stream is not directed to a terminal.

--no-recurse-submodules
--recurse-submodules=check|on-demand|only|no
May be used to make sure all submodule commits used by the revisions to be pushed are available on a remote-tracking branch. If check is used Git will verify that all submodule commits that changed in the revisions to be pushed are available on at least one remote of the submodule. If any commits are missing the push will be aborted and exit with non-zero status. If on-demand is used all submodules that changed in the revisions to be pushed will be pushed. If on-demand was not able to push all necessary revisions it will also be aborted and exit with non-zero status. If only is used all submodules will be recursively pushed while the superproject is left unpushed. A value of no or using --no-recurse-submodules can be used to override the push.recurseSubmodules configuration variable when no submodule recursion is required.

--no-verify
Toggle the pre-push hook (see githooks[5]). The default is --verify, giving the hook a chance to prevent the push. With --no-verify, the hook is bypassed completely.

-4
--ipv4
Use IPv4 addresses only, ignoring IPv6 addresses.

-6
--ipv6
Use IPv6 addresses only, ignoring IPv4 addresses.
`

const logFlagsStr = `--follow
Continue listing the history of a file beyond renames (works only for a single file).

--no-decorate
--decorate[=short|full|auto|no]
Print out the ref names of any commits that are shown. If short is specified, the ref name prefixes refs/heads/, refs/tags/ and refs/remotes/ will not be printed. If full is specified, the full ref name (including prefix) will be printed. If auto is specified, then if the output is going to a terminal, the ref names are shown as if short were given, otherwise no ref names are shown. The default option is short.

--decorate-refs=<pattern>
--decorate-refs-exclude=<pattern>
If no --decorate-refs is given, pretend as if all refs were included. For each candidate, do not use it for decoration if it matches any patterns given to --decorate-refs-exclude or if it doesn’t match any of the patterns given to --decorate-refs. The log.excludeDecoration config option allows excluding refs from the decorations, but an explicit --decorate-refs pattern will override a match in log.excludeDecoration.

--source
Print out the ref name given on the command line by which each commit was reached.

--[no-]mailmap
--[no-]use-mailmap
Use mailmap file to map author and committer names and email addresses to canonical real names and email addresses. See git-shortlog[1].

--full-diff
Without this flag, git log -p <path>... shows commits that touch the specified paths, and diffs about the same specified paths. With this, the full diff is shown for commits that touch the specified paths; this means that "<path>…​" limits only commits, and doesn’t limit diff for those commits.

Note that this affects all diff-based output types, e.g. those produced by --stat, etc.

--log-size
Include a line “log size <number>” in the output for each commit, where <number> is the length of that commit’s message in bytes. Intended to speed up tools that read log messages from git log output by allowing them to allocate space in advance.

-L <start>,<end>:<file>
-L :<funcname>:<file>
Trace the evolution of the line range given by "<start>,<end>" (or the function name regex <funcname>) within the <file>. You may not give any pathspec limiters. This is currently limited to a walk starting from a single revision, i.e., you may only give zero or one positive revision arguments, and <start> and <end> (or <funcname>) must exist in the starting revision. You can specify this option more than once. Implies --patch. Patch output can be suppressed using --no-patch, but other diff formats (namely --raw, --numstat, --shortstat, --dirstat, --summary, --name-only, --name-status, --check) are not currently implemented.

<start> and <end> can take one of these forms:

number

If <start> or <end> is a number, it specifies an absolute line number (lines count from 1).

/regex/

This form will use the first line matching the given POSIX regex. If <start> is a regex, it will search from the end of the previous -L range, if any, otherwise from the start of file. If <start> is “^/regex/”, it will search from the start of file. If <end> is a regex, it will search starting at the line given by <start>.

+offset or -offset

This is only valid for <end> and will specify a number of lines before or after the line given by <start>.

If “:<funcname>” is given in place of <start> and <end>, it is a regular expression that denotes the range from the first funcname line that matches <funcname>, up to the next funcname line. “:<funcname>” searches from the end of the previous -L range, if any, otherwise from the start of file. “^:<funcname>” searches from the start of file.

<revision range>
Show only commits in the specified revision range. When no <revision range> is specified, it defaults to HEAD (i.e. the whole history leading to the current commit). origin..HEAD specifies all the commits reachable from the current commit (i.e. HEAD), but not from origin. For a complete list of ways to spell <revision range>, see the Specifying Ranges section of gitrevisions[7].

[--] <path>…​
Show only commits that are enough to explain how the files that match the specified paths came to be. See History Simplification below for details and other simplification modes.

Paths may need to be prefixed with -- to separate them from options or the revision range, when confusion arises.

Commit Limiting
Besides specifying a range of commits that should be listed using the special notations explained in the description, additional commit limiting may be applied.

Using more options generally further limits the output (e.g. --since=<date1> limits to commits newer than <date1>, and using it with --grep=<pattern> further limits to commits whose log message has a line that matches <pattern>), unless otherwise noted.

Note that these are applied before commit ordering and formatting options, such as --reverse.

-<number>
-n <number>
--max-count=<number>
Limit the number of commits to output.

--skip=<number>
Skip number commits before starting to show the commit output.

--since=<date>
--after=<date>
Show commits more recent than a specific date.

--until=<date>
--before=<date>
Show commits older than a specific date.

--author=<pattern>
--committer=<pattern>
Limit the commits output to ones with author/committer header lines that match the specified pattern (regular expression). With more than one --author=<pattern>, commits whose author matches any of the given patterns are chosen (similarly for multiple --committer=<pattern>).

--grep-reflog=<pattern>
Limit the commits output to ones with reflog entries that match the specified pattern (regular expression). With more than one --grep-reflog, commits whose reflog message matches any of the given patterns are chosen. It is an error to use this option unless --walk-reflogs is in use.

--grep=<pattern>
Limit the commits output to ones with log message that matches the specified pattern (regular expression). With more than one --grep=<pattern>, commits whose message matches any of the given patterns are chosen (but see --all-match).

When --notes is in effect, the message from the notes is matched as if it were part of the log message.

--all-match
Limit the commits output to ones that match all given --grep, instead of ones that match at least one.

--invert-grep
Limit the commits output to ones with log message that do not match the pattern specified with --grep=<pattern>.

-i
--regexp-ignore-case
Match the regular expression limiting patterns without regard to letter case.

--basic-regexp
Consider the limiting patterns to be basic regular expressions; this is the default.

-E
--extended-regexp
Consider the limiting patterns to be extended regular expressions instead of the default basic regular expressions.

-F
--fixed-strings
Consider the limiting patterns to be fixed strings (don’t interpret pattern as a regular expression).

-P
--perl-regexp
Consider the limiting patterns to be Perl-compatible regular expressions.

Support for these types of regular expressions is an optional compile-time dependency. If Git wasn’t compiled with support for them providing this option will cause it to die.

--remove-empty
Stop when a given path disappears from the tree.

--merges
Print only merge commits. This is exactly the same as --min-parents=2.

--no-merges
Do not print commits with more than one parent. This is exactly the same as --max-parents=1.

--min-parents=<number>
--max-parents=<number>
--no-min-parents
--no-max-parents
Show only commits which have at least (or at most) that many parent commits. In particular, --max-parents=1 is the same as --no-merges, --min-parents=2 is the same as --merges. --max-parents=0 gives all root commits and --min-parents=3 all octopus merges.

--no-min-parents and --no-max-parents reset these limits (to no limit) again. Equivalent forms are --min-parents=0 (any commit has 0 or more parents) and --max-parents=-1 (negative numbers denote no upper limit).

--first-parent
Follow only the first parent commit upon seeing a merge commit. This option can give a better overview when viewing the evolution of a particular topic branch, because merges into a topic branch tend to be only about adjusting to updated upstream from time to time, and this option allows you to ignore the individual commits brought in to your history by such a merge. Cannot be combined with --bisect.

--not
Reverses the meaning of the ^ prefix (or lack thereof) for all following revision specifiers, up to the next --not.

--all
Pretend as if all the refs in refs/, along with HEAD, are listed on the command line as <commit>.

--branches[=<pattern>]
Pretend as if all the refs in refs/heads are listed on the command line as <commit>. If <pattern> is given, limit branches to ones matching given shell glob. If pattern lacks ?, *, or [, /* at the end is implied.

--tags[=<pattern>]
Pretend as if all the refs in refs/tags are listed on the command line as <commit>. If <pattern> is given, limit tags to ones matching given shell glob. If pattern lacks ?, *, or [, /* at the end is implied.

--remotes[=<pattern>]
Pretend as if all the refs in refs/remotes are listed on the command line as <commit>. If <pattern> is given, limit remote-tracking branches to ones matching given shell glob. If pattern lacks ?, *, or [, /* at the end is implied.

--glob=<glob-pattern>
Pretend as if all the refs matching shell glob <glob-pattern> are listed on the command line as <commit>. Leading refs/, is automatically prepended if missing. If pattern lacks ?, *, or [, /* at the end is implied.

--exclude=<glob-pattern>
Do not include refs matching <glob-pattern> that the next --all, --branches, --tags, --remotes, or --glob would otherwise consider. Repetitions of this option accumulate exclusion patterns up to the next --all, --branches, --tags, --remotes, or --glob option (other options or arguments do not clear accumulated patterns).

The patterns given should not begin with refs/heads, refs/tags, or refs/remotes when applied to --branches, --tags, or --remotes, respectively, and they must begin with refs/ when applied to --glob or --all. If a trailing /* is intended, it must be given explicitly.

--reflog
Pretend as if all objects mentioned by reflogs are listed on the command line as <commit>.

--alternate-refs
Pretend as if all objects mentioned as ref tips of alternate repositories were listed on the command line. An alternate repository is any repository whose object directory is specified in objects/info/alternates. The set of included objects may be modified by core.alternateRefsCommand, etc. See git-config[1].

--single-worktree
By default, all working trees will be examined by the following options when there are more than one (see git-worktree[1]): --all, --reflog and --indexed-objects. This option forces them to examine the current working tree only.

--ignore-missing
Upon seeing an invalid object name in the input, pretend as if the bad input was not given.

--bisect
Pretend as if the bad bisection ref refs/bisect/bad was listed and as if it was followed by --not and the good bisection refs refs/bisect/good-* on the command line. Cannot be combined with --first-parent.

--stdin
In addition to the <commit> listed on the command line, read them from the standard input. If a -- separator is seen, stop reading commits and start reading paths to limit the result.

--cherry-mark
Like --cherry-pick (see below) but mark equivalent commits with = rather than omitting them, and inequivalent ones with +.

--cherry-pick
Omit any commit that introduces the same change as another commit on the “other side” when the set of commits are limited with symmetric difference.

For example, if you have two branches, A and B, a usual way to list all commits on only one side of them is with --left-right (see the example below in the description of the --left-right option). However, it shows the commits that were cherry-picked from the other branch (for example, “3rd on b” may be cherry-picked from branch A). With this option, such pairs of commits are excluded from the output.

--left-only
--right-only
List only commits on the respective side of a symmetric difference, i.e. only those which would be marked < resp. > by --left-right.

For example, --cherry-pick --right-only A...B omits those commits from B which are in A or are patch-equivalent to a commit in A. In other words, this lists the + commits from git cherry A B. More precisely, --cherry-pick --right-only --no-merges gives the exact list.

--cherry
A synonym for --right-only --cherry-mark --no-merges; useful to limit the output to the commits on our side and mark those that have been applied to the other side of a forked history with git log --cherry upstream...mybranch, similar to git cherry upstream mybranch.

-g
--walk-reflogs
Instead of walking the commit ancestry chain, walk reflog entries from the most recent one to older ones. When this option is used you cannot specify commits to exclude (that is, ^commit, commit1..commit2, and commit1...commit2 notations cannot be used).

With --pretty format other than oneline and reference (for obvious reasons), this causes the output to have two extra lines of information taken from the reflog. The reflog designator in the output may be shown as ref@{Nth} (where Nth is the reverse-chronological index in the reflog) or as ref@{timestamp} (with the timestamp for that entry), depending on a few rules:

If the starting point is specified as ref@{Nth}, show the index format.

If the starting point was specified as ref@{now}, show the timestamp format.

If neither was used, but --date was given on the command line, show the timestamp in the format requested by --date.

Otherwise, show the index format.

Under --pretty=oneline, the commit message is prefixed with this information on the same line. This option cannot be combined with --reverse. See also git-reflog[1].

Under --pretty=reference, this information will not be shown at all.

--merge
After a failed merge, show refs that touch files having a conflict and don’t exist on all heads to merge.

--boundary
Output excluded boundary commits. Boundary commits are prefixed with -.

History Simplification
Sometimes you are only interested in parts of the history, for example the commits modifying a particular <path>. But there are two parts of History Simplification, one part is selecting the commits and the other is how to do it, as there are various strategies to simplify the history.

The following options select the commits to be shown:

<paths>
Commits modifying the given <paths> are selected.

--simplify-by-decoration
Commits that are referred by some branch or tag are selected.

Note that extra commits can be shown to give a meaningful history.

The following options affect the way the simplification is performed:

Default mode
Simplifies the history to the simplest history explaining the final state of the tree. Simplest because it prunes some side branches if the end result is the same (i.e. merging branches with the same content)

--show-pulls
Include all commits from the default mode, but also any merge commits that are not TREESAME to the first parent but are TREESAME to a later parent. This mode is helpful for showing the merge commits that "first introduced" a change to a branch.

--full-history
Same as the default mode, but does not prune some history.

--dense
Only the selected commits are shown, plus some to have a meaningful history.

--sparse
All commits in the simplified history are shown.

--simplify-merges
Additional option to --full-history to remove some needless merges from the resulting history, as there are no selected commits contributing to this merge.

--ancestry-path
When given a range of commits to display (e.g. commit1..commit2 or commit2 ^commit1), only display commits that exist directly on the ancestry chain between the commit1 and commit2, i.e. commits that are both descendants of commit1, and ancestors of commit2.

A more detailed explanation follows.

Suppose you specified foo as the <paths>. We shall call commits that modify foo !TREESAME, and the rest TREESAME. (In a diff filtered for foo, they look different and equal, respectively.)

In the following, we will always refer to the same example history to illustrate the differences between simplification settings. We assume that you are filtering for a file foo in this commit graph:

	  .-A---M---N---O---P---Q
	 /     /   /   /   /   /
	I     B   C   D   E   Y
	 \   /   /   /   /   /
	  -------------   X
The horizontal line of history A---Q is taken to be the first parent of each merge. The commits are:

I is the initial commit, in which foo exists with contents “asdf”, and a file quux exists with contents “quux”. Initial commits are compared to an empty tree, so I is !TREESAME.

In A, foo contains just “foo”.

B contains the same change as A. Its merge M is trivial and hence TREESAME to all parents.

C does not change foo, but its merge N changes it to “foobar”, so it is not TREESAME to any parent.

D sets foo to “baz”. Its merge O combines the strings from N and D to “foobarbaz”; i.e., it is not TREESAME to any parent.

E changes quux to “xyzzy”, and its merge P combines the strings to “quux xyzzy”. P is TREESAME to O, but not to E.

X is an independent root commit that added a new file side, and Y modified it. Y is TREESAME to X. Its merge Q added side to P, and Q is TREESAME to P, but not to Y.

rev-list walks backwards through history, including or excluding commits based on whether --full-history and/or parent rewriting (via --parents or --children) are used. The following settings are available.

Default mode
Commits are included if they are not TREESAME to any parent (though this can be changed, see --sparse below). If the commit was a merge, and it was TREESAME to one parent, follow only that parent. (Even if there are several TREESAME parents, follow only one of them.) Otherwise, follow all parents.

This results in:

.-A---N---O
/     /   /
I---------D
Note how the rule to only follow the TREESAME parent, if one is available, removed B from consideration entirely. C was considered via N, but is TREESAME. Root commits are compared to an empty tree, so I is !TREESAME.

Parent/child relations are only visible with --parents, but that does not affect the commits selected in default mode, so we have shown the parent lines.

--full-history without parent rewriting
This mode differs from the default in one point: always follow all parents of a merge, even if it is TREESAME to one of them. Even if more than one side of the merge has commits that are included, this does not imply that the merge itself is! In the example, we get

I  A  B  N  D  O  P  Q
M was excluded because it is TREESAME to both parents. E, C and B were all walked, but only B was !TREESAME, so the others do not appear.

Note that without parent rewriting, it is not really possible to talk about the parent/child relationships between the commits, so we show them disconnected.

--full-history with parent rewriting
Ordinary commits are only included if they are !TREESAME (though this can be changed, see --sparse below).

Compare to --full-history without rewriting above. Note that E was pruned away because it is TREESAME, but the parent list of P was rewritten to contain Es parent I. The same happened for C and N, and X, Y and Q.

In addition to the above settings, you can change whether TREESAME affects inclusion:

--dense
Commits that are walked are included if they are not TREESAME to any parent.

--sparse
All commits that are walked are included.

Note that without --full-history, this still simplifies merges: if one of the parents is TREESAME, we follow only that one, so the other sides of the merge are never walked.

--simplify-merges
First, build a history graph in the same way that --full-history with parent rewriting does (see above).

Then simplify each commit C to its replacement C Prime in the final history according to the following rules:

Set C Prime to C.

Replace each parent P of C Prime with its simplification P Prime. In the process, drop parents that are ancestors of other parents or that are root commits TREESAME to an empty tree, and remove duplicates, but take care to never drop all parents that we are TREESAME to.

If after this parent rewriting, C' is a root or merge commit (has zero or >1 parents), a boundary commit, or !TREESAME, it remains. Otherwise, it is replaced with its only parent.

The effect of this is best shown by way of comparing to --full-history with parent rewriting. The example turns into:

	  .-A---M---N---O
	 /     /       /
	I     B       D
	 \   /       /
	  ---------
Note the major differences in N, P, and Q over --full-history:

N's parent list had I removed, because it is an ancestor of the other parent M. Still, N remained because it is !TREESAME.

P's parent list similarly had I removed. P was then removed completely, because it had one parent and is TREESAME.

Q's parent list had Y simplified to X. X was then removed, because it was a TREESAME root. Q was then removed completely, because it had one parent and is TREESAME.

There is another simplification mode available:

--ancestry-path
Limit the displayed commits to those directly on the ancestry chain between the “from” and “to” commits in the given commit range. I.e. only display commits that are ancestor of the “to” commit and descendants of the “from” commit.

As an example use case, consider the following commit history:

D---E-------F
/     \       \
B---C---G---H---I---J
/                     \
A-------K---------------L--M
A regular D..M computes the set of commits that are ancestors of M, but excludes the ones that are ancestors of D. This is useful to see what happened to the history leading to M since D, in the sense that “what does M have that did not exist in D”. The result in this example would be all the commits, except A and B (and D itself, of course).

When we want to find out what commits in M are contaminated with the bug introduced by D and need fixing, however, we might want to view only the subset of D..M that are actually descendants of D, i.e. excluding C and K. This is exactly what the --ancestry-path option does. Applied to the D..M range, it results in:

E-------F
\       \
G---H---I---J
\
L--M
Before discussing another option, --show-pulls, we need to create a new example history.

A common problem users face when looking at simplified history is that a commit they know changed a file somehow does not appear in the file’s simplified history. Let’s demonstrate a new example and show how options such as --full-history and --simplify-merges works in that case:

--show-pulls
In addition to the commits shown in the default history, show each merge commit that is not TREESAME to its first parent but is TREESAME to a later parent.

When a merge commit is included by --show-pulls, the merge is treated as if it "pulled" the change from another branch. When using --show-pulls on this example (and no other options) the resulting graph is:

	I---X---R---N
Here, the merge commits R and N are included because they pulled the commits X and R into the base branch, respectively. These merges are the reason the commits A and B do not appear in the default history.

When --show-pulls is paired with --simplify-merges, the graph includes all of the necessary information:

The --simplify-by-decoration option allows you to view only the big picture of the topology of the history, by omitting commits that are not referenced by tags. Commits are marked as !TREESAME (in other words, kept after history simplification rules described above) if (1) they are referenced by tags, or (2) they change the contents of the paths given on the command line. All other commits are marked as TREESAME (subject to be simplified away).

Commit Ordering
By default, the commits are shown in reverse chronological order.

--date-order
Show no parents before all of its children are shown, but otherwise show commits in the commit timestamp order.

--author-date-order
Show no parents before all of its children are shown, but otherwise show commits in the author timestamp order.

--topo-order
Show no parents before all of its children are shown, and avoid showing commits on multiple lines of history intermixed.

For example, in a commit history like this:

---1----2----4----7
\	       \
3----5----6----8---
where the numbers denote the order of commit timestamps, git rev-list and friends with --date-order show the commits in the timestamp order: 8 7 6 5 4 3 2 1.

With --topo-order, they would show 8 6 5 3 7 4 2 1 (or 8 7 4 2 6 5 3 1); some older commits are shown before newer ones in order to avoid showing the commits from two parallel development track mixed together.

--reverse
Output the commits chosen to be shown (see Commit Limiting section above) in reverse order. Cannot be combined with --walk-reflogs.

Object Traversal
These options are mostly targeted for packing of Git repositories.

--no-walk[=(sorted|unsorted)]
Only show the given commits, but do not traverse their ancestors. This has no effect if a range is specified. If the argument unsorted is given, the commits are shown in the order they were given on the command line. Otherwise (if sorted or no argument was given), the commits are shown in reverse chronological order by commit time. Cannot be combined with --graph.

--do-walk
Overrides a previous --no-walk.

Commit Formatting
--pretty[=<format>]
--format=<format>
Pretty-print the contents of the commit logs in a given format, where <format> can be one of oneline, short, medium, full, fuller, reference, email, raw, format:<string> and tformat:<string>. When <format> is none of the above, and has %placeholder in it, it acts as if --pretty=tformat:<format> were given.

See the "PRETTY FORMATS" section for some additional details for each format. When =<format> part is omitted, it defaults to medium.

Note: you can specify the default pretty format in the repository configuration (see git-config[1]).

--abbrev-commit
Instead of showing the full 40-byte hexadecimal commit object name, show only a partial prefix. Non default number of digits can be specified with "--abbrev=<n>" (which also modifies diff output, if it is displayed).

This should make "--pretty=oneline" a whole lot more readable for people using 80-column terminals.

--no-abbrev-commit
Show the full 40-byte hexadecimal commit object name. This negates --abbrev-commit and those options which imply it such as "--oneline". It also overrides the log.abbrevCommit variable.

--oneline
This is a shorthand for "--pretty=oneline --abbrev-commit" used together.

--encoding=<encoding>
The commit objects record the encoding used for the log message in their encoding header; this option can be used to tell the command to re-code the commit log message in the encoding preferred by the user. For non plumbing commands this defaults to UTF-8. Note that if an object claims to be encoded in X and we are outputting in X, we will output the object verbatim; this means that invalid sequences in the original commit may be copied to the output.

--expand-tabs=<n>
--expand-tabs
--no-expand-tabs
Perform a tab expansion (replace each tab with enough spaces to fill to the next display column that is multiple of <n>) in the log message before showing it in the output. --expand-tabs is a short-hand for --expand-tabs=8, and --no-expand-tabs is a short-hand for --expand-tabs=0, which disables tab expansion.

By default, tabs are expanded in pretty formats that indent the log message by 4 spaces (i.e. medium, which is the default, full, and fuller).

--notes[=<ref>]
Show the notes (see git-notes[1]) that annotate the commit, when showing the commit log message. This is the default for git log, git show and git whatchanged commands when there is no --pretty, --format, or --oneline option given on the command line.

By default, the notes shown are from the notes refs listed in the core.notesRef and notes.displayRef variables (or corresponding environment overrides). See git-config[1] for more details.

With an optional <ref> argument, use the ref to find the notes to display. The ref can specify the full refname when it begins with refs/notes/; when it begins with notes/, refs/ and otherwise refs/notes/ is prefixed to form a full name of the ref.

Multiple --notes options can be combined to control which notes are being displayed. Examples: "--notes=foo" will show only notes from "refs/notes/foo"; "--notes=foo --notes" will show both notes from "refs/notes/foo" and from the default notes ref(s).

--no-notes
Do not show notes. This negates the above --notes option, by resetting the list of notes refs from which notes are shown. Options are parsed in the order given on the command line, so e.g. "--notes --notes=foo --no-notes --notes=bar" will only show notes from "refs/notes/bar".

--show-notes[=<ref>]
--[no-]standard-notes
These options are deprecated. Use the above --notes/--no-notes options instead.

--show-signature
Check the validity of a signed commit object by passing the signature to gpg --verify and show the output.

--relative-date
Synonym for --date=relative.

--date=<format>
Only takes effect for dates shown in human-readable format, such as when using --pretty. log.date config variable sets a default value for the log command’s --date option. By default, dates are shown in the original time zone (either committer’s or author’s). If -local is appended to the format (e.g., iso-local), the user’s local time zone is used instead.

--date=relative shows dates relative to the current time, e.g. “2 hours ago”. The -local option has no effect for --date=relative.

--date=local is an alias for --date=default-local.

--date=iso 
shows timestamps in a ISO 8601-like format. The differences to the strict ISO 8601 format are:

a space instead of the T date/time delimiter

a space between time and time zone

no colon between hours and minutes of the time zone

--date=iso-strict  shows timestamps in strict ISO 8601 format.

--date=rfc  shows timestamps in RFC 2822 format, often found in email messages.

--date=short shows only the date, but not the time, in YYYY-MM-DD format.

--date=raw shows the date as seconds since the epoch (1970-01-01 00:00:00 UTC), followed by a space, and then the timezone as an offset from UTC (a + or - with four digits; the first two are hours, and the second two are minutes). I.e., as if the timestamp were formatted with strftime("%s %z")). Note that the -local option does not affect the seconds-since-epoch value (which is always measured in UTC), but does switch the accompanying timezone value.

--date=human shows the timezone if the timezone does not match the current time-zone, and doesn’t print the whole date if that matches (ie skip printing year for dates that are "this year", but also skip the whole date itself if it’s in the last few days and we can just say what weekday it was). For older dates the hour and minute is also omitted.

--date=unix shows the date as a Unix epoch timestamp (seconds since 1970). As with --raw, this is always in UTC and therefore -local has no effect.

--date=format:... feeds the format ... to your system strftime, except for %z and %Z, which are handled internally. Use --date=format:%c to show the date in your system locale’s preferred format. See the strftime manual for a complete list of format placeholders. When using -local, the correct syntax is --date=format-local:....

--date=default 
is the default format, and is similar to --date=rfc2822, with a few exceptions:

there is no comma after the day-of-week

the time zone is omitted when the local time zone is used

--parents
Print also the parents of the commit (in the form "commit parent…​"). Also enables parent rewriting, see History Simplification above.

--children
Print also the children of the commit (in the form "commit child…​"). Also enables parent rewriting, see History Simplification above.

--left-right
Mark which side of a symmetric difference a commit is reachable from. Commits from the left side are prefixed with < and those from the right with >. If combined with --boundary, those commits are prefixed with -.

For example, if you have this topology:

y---b---b  branch B
/ \ /
/   .
/   / \
o---x---a---a  branch A
you would get an output like this:

$ git rev-list --left-right --boundary --pretty=oneline A...B

>bbbbbbb... 3rd on b
>bbbbbbb... 2nd on b
<aaaaaaa... 3rd on a
<aaaaaaa... 2nd on a
-yyyyyyy... 1st on b
-xxxxxxx... 1st on a
--graph
Draw a text-based graphical representation of the commit history on the left hand side of the output. This may cause extra lines to be printed in between commits, in order for the graph history to be drawn properly. Cannot be combined with --no-walk.

This enables parent rewriting, see History Simplification above.

This implies the --topo-order option by default, but the --date-order option may also be specified.

--show-linear-break[=<barrier>]
When --graph is not used, all history branches are flattened which can make it hard to see that the two consecutive commits do not belong to a linear branch. This option puts a barrier in between them in that case. If <barrier> is specified, it is the string that will be shown instead of the default one.

Diff Formatting
Listed below are options that control the formatting of diff output. Some of them are specific to git-rev-list[1], however other diff options may be given. See git-diff-files[1] for more options.

-c
With this option, diff output for a merge commit shows the differences from each of the parents to the merge result simultaneously instead of showing pairwise diff between a parent and the result one at a time. Furthermore, it lists only files which were modified from all parents.

--cc
This flag implies the -c option and further compresses the patch output by omitting uninteresting hunks whose contents in the parents have only two variants and the merge result picks one of them without modification.

--combined-all-paths
This flag causes combined diffs (used for merge commits) to list the name of the file from all parents. It thus only has effect when -c or --cc are specified, and is likely only useful if filename changes are detected (i.e. when either rename or copy detection have been requested).

-m
This flag makes the merge commits show the full diff like regular commits; for each merge parent, a separate log entry and diff is generated. An exception is that only diff against the first parent is shown when --first-parent option is given; in that case, the output represents the changes the merge brought into the then-current branch.

-r
Show recursive diffs.

-t
Show the tree objects in the diff output. This implies -r.`

const rebaseFlagsStr = `--onto <newbase>
Starting point at which to create the new commits. If the --onto option is not specified, the starting point is <upstream>. May be any valid commit, and not just an existing branch name.

As a special case, you may use "A...B" as a shortcut for the merge base of A and B if there is exactly one merge base. You can leave out at most one of A and B, in which case it defaults to HEAD.

--keep-base
Set the starting point at which to create the new commits to the merge base of <upstream> <branch>. Running git rebase --keep-base <upstream> <branch> is equivalent to running git rebase --onto <upstream>…​ <upstream>.

This option is useful in the case where one is developing a feature on top of an upstream branch. While the feature is being worked on, the upstream branch may advance and it may not be the best idea to keep rebasing on top of the upstream but to keep the base commit as-is.

Although both this option and --fork-point find the merge base between <upstream> and <branch>, this option uses the merge base as the starting point on which new commits will be created, whereas --fork-point uses the merge base to determine the set of commits which will be rebased.

See also INCOMPATIBLE OPTIONS below.

<upstream>
Upstream branch to compare against. May be any valid commit, not just an existing branch name. Defaults to the configured upstream for the current branch.

<branch>
Working branch; defaults to HEAD.

--continue
Restart the rebasing process after having resolved a merge conflict.

--abort
Abort the rebase operation and reset HEAD to the original branch. If <branch> was provided when the rebase operation was started, then HEAD will be reset to <branch>. Otherwise HEAD will be reset to where it was when the rebase operation was started.

--quit
Abort the rebase operation but HEAD is not reset back to the original branch. The index and working tree are also left unchanged as a result. If a temporary stash entry was created using --autostash, it will be saved to the stash list.

--apply
Use applying strategies to rebase (calling git-am internally). This option may become a no-op in the future once the merge backend handles everything the apply one does.

See also INCOMPATIBLE OPTIONS below.

--empty={drop,keep,ask}
How to handle commits that are not empty to start and are not clean cherry-picks of any upstream commit, but which become empty after rebasing (because they contain a subset of already upstream changes). With drop (the default), commits that become empty are dropped. With keep, such commits are kept. With ask (implied by --interactive), the rebase will halt when an empty commit is applied allowing you to choose whether to drop it, edit files more, or just commit the empty changes. Other options, like --exec, will use the default of drop unless -i/--interactive is explicitly specified.

Note that commits which start empty are kept (unless --no-keep-empty is specified), and commits which are clean cherry-picks (as determined by git log --cherry-mark ...) are detected and dropped as a preliminary step (unless --reapply-cherry-picks is passed).

See also INCOMPATIBLE OPTIONS below.

--no-keep-empty
--keep-empty
Do not keep commits that start empty before the rebase (i.e. that do not change anything from its parent) in the result. The default is to keep commits which start empty, since creating such commits requires passing the --allow-empty override flag to git commit, signifying that a user is very intentionally creating such a commit and thus wants to keep it.

Usage of this flag will probably be rare, since you can get rid of commits that start empty by just firing up an interactive rebase and removing the lines corresponding to the commits you don’t want. This flag exists as a convenient shortcut, such as for cases where external tools generate many empty commits and you want them all removed.

For commits which do not start empty but become empty after rebasing, see the --empty flag.

See also INCOMPATIBLE OPTIONS below.

--reapply-cherry-picks
--no-reapply-cherry-picks
Reapply all clean cherry-picks of any upstream commit instead of preemptively dropping them. (If these commits then become empty after rebasing, because they contain a subset of already upstream changes, the behavior towards them is controlled by the --empty flag.)

By default (or if --no-reapply-cherry-picks is given), these commits will be automatically dropped. Because this necessitates reading all upstream commits, this can be expensive in repos with a large number of upstream commits that need to be read.

--reapply-cherry-picks allows rebase to forgo reading all upstream commits, potentially improving performance.

See also INCOMPATIBLE OPTIONS below.

--allow-empty-message
No-op. Rebasing commits with an empty message used to fail and this option would override that behavior, allowing commits with empty messages to be rebased. Now commits with an empty message do not cause rebasing to halt.

See also INCOMPATIBLE OPTIONS below.

--skip
Restart the rebasing process by skipping the current patch.

--edit-todo
Edit the todo list during an interactive rebase.

--show-current-patch
Show the current patch in an interactive rebase or when rebase is stopped because of conflicts. This is the equivalent of git show REBASE_HEAD.

-m
--merge
Use merging strategies to rebase. When the recursive (default) merge strategy is used, this allows rebase to be aware of renames on the upstream side. This is the default.

Note that a rebase merge works by replaying each commit from the working branch on top of the <upstream> branch. Because of this, when a merge conflict happens, the side reported as ours is the so-far rebased series, starting with <upstream>, and theirs is the working branch. In other words, the sides are swapped.

See also INCOMPATIBLE OPTIONS below.

-s <strategy>
--strategy=<strategy>
Use the given merge strategy. If there is no -s option git merge-recursive is used instead. This implies --merge.

Because git rebase replays each commit from the working branch on top of the <upstream> branch using the given strategy, using the ours strategy simply empties all patches from the <branch>, which makes little sense.

See also INCOMPATIBLE OPTIONS below.

-X <strategy-option>
--strategy-option=<strategy-option>
Pass the <strategy-option> through to the merge strategy. This implies --merge and, if no strategy has been specified, -s recursive. Note the reversal of ours and theirs as noted above for the -m option.

See also INCOMPATIBLE OPTIONS below.

--rerere-autoupdate
--no-rerere-autoupdate
Allow the rerere mechanism to update the index with the result of auto-conflict resolution if possible.

-S[<keyid>]
--gpg-sign[=<keyid>]
--no-gpg-sign
GPG-sign commits. The keyid argument is optional and defaults to the committer identity; if specified, it must be stuck to the option without a space. --no-gpg-sign is useful to countermand both commit.gpgSign configuration variable, and earlier --gpg-sign.

-q
--quiet
Be quiet. Implies --no-stat.

-v
--verbose
Be verbose. Implies --stat.

--stat
Show a diffstat of what changed upstream since the last rebase. The diffstat is also controlled by the configuration option rebase.stat.

-n
--no-stat
Do not show a diffstat as part of the rebase process.

--no-verify
This option bypasses the pre-rebase hook. See also githooks[5].

--verify
Allows the pre-rebase hook to run, which is the default. This option can be used to override --no-verify. See also githooks[5].

-C<n>
Ensure at least <n> lines of surrounding context match before and after each change. When fewer lines of surrounding context exist they all must match. By default no context is ever ignored. Implies --apply.

See also INCOMPATIBLE OPTIONS below.

--no-ff
--force-rebase
-f
Individually replay all rebased commits instead of fast-forwarding over the unchanged ones. This ensures that the entire history of the rebased branch is composed of new commits.

You may find this helpful after reverting a topic branch merge, as this option recreates the topic branch with fresh commits so it can be remerged successfully without needing to "revert the reversion" (see the revert-a-faulty-merge How-To for details).

--fork-point
--no-fork-point
Use reflog to find a better common ancestor between <upstream> and <branch> when calculating which commits have been introduced by <branch>.

When --fork-point is active, fork_point will be used instead of <upstream> to calculate the set of commits to rebase, where fork_point is the result of git merge-base --fork-point <upstream> <branch> command (see git-merge-base[1]). If fork_point ends up being empty, the <upstream> will be used as a fallback.

If <upstream> is given on the command line, then the default is --no-fork-point, otherwise the default is --fork-point.

If your branch was based on <upstream> but <upstream> was rewound and your branch contains commits which were dropped, this option can be used with --keep-base in order to drop those commits from your branch.

See also INCOMPATIBLE OPTIONS below.

--ignore-whitespace
--whitespace=<option>
These flags are passed to the git apply program (see git-apply[1]) that applies the patch. Implies --apply.

See also INCOMPATIBLE OPTIONS below.

--committer-date-is-author-date
--ignore-date
These flags are passed to git am to easily change the dates of the rebased commits (see git-am[1]).

See also INCOMPATIBLE OPTIONS below.

--signoff
Add a Signed-off-by: trailer to all the rebased commits. Note that if --interactive is given then only commits marked to be picked, edited or reworded will have the trailer added.

See also INCOMPATIBLE OPTIONS below.

-i
--interactive
Make a list of the commits which are about to be rebased. Let the user edit that list before rebasing. This mode can also be used to split commits (see SPLITTING COMMITS below).

The commit list format can be changed by setting the configuration option rebase.instructionFormat. A customized instruction format will automatically have the long commit hash prepended to the format.

See also INCOMPATIBLE OPTIONS below.

-r
--rebase-merges[=(rebase-cousins|no-rebase-cousins)]
By default, a rebase will simply drop merge commits from the todo list, and put the rebased commits into a single, linear branch. With --rebase-merges, the rebase will instead try to preserve the branching structure within the commits that are to be rebased, by recreating the merge commits. Any resolved merge conflicts or manual amendments in these merge commits will have to be resolved/re-applied manually.

By default, or when no-rebase-cousins was specified, commits which do not have <upstream> as direct ancestor will keep their original branch point, i.e. commits that would be excluded by git-log[1]'s --ancestry-path option will keep their original ancestry by default. If the rebase-cousins mode is turned on, such commits are instead rebased onto <upstream> (or <onto>, if specified).

The --rebase-merges mode is similar in spirit to the deprecated --preserve-merges but works with interactive rebases, where commits can be reordered, inserted and dropped at will.

It is currently only possible to recreate the merge commits using the recursive merge strategy; Different merge strategies can be used only via explicit exec git merge -s <strategy> [...] commands.

See also REBASING MERGES and INCOMPATIBLE OPTIONS below.

-p
--preserve-merges
[DEPRECATED: use --rebase-merges instead] Recreate merge commits instead of flattening the history by replaying commits a merge commit introduces. Merge conflict resolutions or manual amendments to merge commits are not preserved.

This uses the --interactive machinery internally, but combining it with the --interactive option explicitly is generally not a good idea unless you know what you are doing (see BUGS below).

See also INCOMPATIBLE OPTIONS below.

-x <cmd>
--exec <cmd>
Append "exec <cmd>" after each line creating a commit in the final history. <cmd> will be interpreted as one or more shell commands. Any command that fails will interrupt the rebase, with exit code 1.

You may execute several commands by either using one instance of --exec with several commands:

git rebase -i --exec "cmd1 && cmd2 && ..."
or by giving more than one --exec:

git rebase -i --exec "cmd1" --exec "cmd2" --exec ...
If --autosquash is used, "exec" lines will not be appended for the intermediate commits, and will only appear at the end of each squash/fixup series.

This uses the --interactive machinery internally, but it can be run without an explicit --interactive.

See also INCOMPATIBLE OPTIONS below.

--root
Rebase all commits reachable from <branch>, instead of limiting them with an <upstream>. This allows you to rebase the root commit(s) on a branch. When used with --onto, it will skip changes already contained in <newbase> (instead of <upstream>) whereas without --onto it will operate on every change. When used together with both --onto and --preserve-merges, all root commits will be rewritten to have <newbase> as parent instead.

See also INCOMPATIBLE OPTIONS below.

--autosquash
--no-autosquash
When the commit log message begins with "squash! …​" (or "fixup! …​"), and there is already a commit in the todo list that matches the same ..., automatically modify the todo list of rebase -i so that the commit marked for squashing comes right after the commit to be modified, and change the action of the moved commit from pick to squash (or fixup). A commit matches the ... if the commit subject matches, or if the ... refers to the commit’s hash. As a fall-back, partial matches of the commit subject work, too. The recommended way to create fixup/squash commits is by using the --fixup/--squash options of git-commit[1].

If the --autosquash option is enabled by default using the configuration variable rebase.autoSquash, this option can be used to override and disable this setting.

See also INCOMPATIBLE OPTIONS below.

--autostash
--no-autostash
Automatically create a temporary stash entry before the operation begins, and apply it after the operation ends. This means that you can run rebase on a dirty worktree. However, use with care: the final stash application after a successful rebase might result in non-trivial conflicts.

--reschedule-failed-exec
--no-reschedule-failed-exec
Automatically reschedule exec commands that failed. This only makes sense in interactive mode (or when an --exec option was provided).`

const resetFlagsStr = `--soft
Does not touch the index file or the working tree at all (but resets the head to <commit>, just like all modes do). This leaves all your changed files "Changes to be committed", as git status would put it.

--mixed
Resets the index but not the working tree (i.e., the changed files are preserved but not marked for commit) and reports what has not been updated. This is the default action.

If -N is specified, removed paths are marked as intent-to-add (see git-add[1]).

--hard
Resets the index and working tree. Any changes to tracked files in the working tree since <commit> are discarded.

--merge
Resets the index and updates the files in the working tree that are different between <commit> and HEAD, but keeps those which are different between the index and working tree (i.e. which have changes which have not been added). If a file that is different between <commit> and the index has unstaged changes, reset is aborted.

In other words, --merge does something like a git read-tree -u -m <commit>, but carries forward unmerged index entries.

--keep
Resets index entries and updates files in the working tree that are different between <commit> and HEAD. If a file that is different between <commit> and HEAD has local changes, reset is aborted.

--[no-]recurse-submodules
When the working tree is updated, using --recurse-submodules will also recursively reset the working tree of all active submodules according to the commit recorded in the superproject, also setting the submodules' HEAD to be detached at that commit.`

package cmd

import (
	"fmt"
	"github.com/chriswalz/complete/v3"
	"github.com/spf13/cobra"
	"github.com/thoas/go-funk"
	"strings"
)

// gitmojiCmd represents the gitmoji command
var gitmojiCmd = &cobra.Command{
	Use:   "gitmoji",
	Short: "(Pre-alpha) Commit using gitmojis",
	Long:  `bit save gitmoji"`,
	Run: func(cmd *cobra.Command, args []string) {
		emojiAndMsg := ""
		if len(args) == 0 {
			gitmojiSuggestions := GitmojiSuggestions()
			suggestionTree := &complete.CompTree{
				Sub: map[string]*complete.CompTree{
					"gitmoji": {
						Dynamic: toAutoCLI(gitmojiSuggestions),
					},
				},
			}
			emojiAndMsg = SuggestionPrompt("> bit gitmoji ", specificCommandCompleter("gitmoji", suggestionTree))
		} else {
			emojiAndMsg = args[0]
			if len(args) > 0 {
				emojiAndMsg = strings.Join(args, " ")
			}
		}

		if len(emojiAndMsg) < 1 {
			fmt.Println("commit message missing")
			return
		}
		g := funk.Find(gitmojis, func(g *Gitmoji) bool {
			if strings.Contains(emojiAndMsg, g.Description) {
				return true
			}
			return false
		})
		if g == nil {
			fmt.Println("No related gitmoji found")
			return
		}
		emojiAndMsgWithoutEmojiDescription := strings.ReplaceAll(emojiAndMsg, g.(*Gitmoji).Description, g.(*Gitmoji).Emoji)
		save([]string{"-m " + emojiAndMsgWithoutEmojiDescription})
	},
}

func init() {
	BitCmd.AddCommand(gitmojiCmd)
}

func GitmojiSuggestions() []complete.Suggestion {
	var suggestions []complete.Suggestion
	for _, gitmoji := range gitmojis {
		suggestions = append(suggestions, complete.Suggestion{
			Name: `"` + gitmoji.Description,
			//Text: "\"" + gitmoji.Emoji + " " + gitmoji.Description,
			//Description: "  " + gitmoji.Emoji + "  " ,
		})
	}
	return suggestions
}

type Gitmoji struct {
	Emoji       string `json:"emoji"`
	Entity      string `json:"entity"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

var gitmojis = []*Gitmoji{
	{
		Emoji:       "🎨",
		Entity:      "&#x1f3a8;",
		Code:        ":art:",
		Description: "Improve structure / format of the code.",
		Name:        "art",
	},
	{
		Emoji:       "⚡️",
		Entity:      "&#x26a1;",
		Code:        ":zap:",
		Description: "Improve performance.",
		Name:        "zap",
	},
	{
		Emoji:       "🔥",
		Entity:      "&#x1f525;",
		Code:        ":fire:",
		Description: "Remove code or files.",
		Name:        "fire",
	},
	{
		Emoji:       "🐛",
		Entity:      "&#x1f41b;",
		Code:        ":bug:",
		Description: "Fix a bug.",
		Name:        "bug",
	},
	{
		Emoji:       "🚑",
		Entity:      "&#128657;",
		Code:        ":ambulance:",
		Description: "Critical hotfix.",
		Name:        "ambulance",
	},
	{
		Emoji:       "✨",
		Entity:      "&#x2728;",
		Code:        ":sparkles:",
		Description: "Introduce new features.",
		Name:        "sparkles",
	},
	{
		Emoji:       "📝",
		Entity:      "&#x1f4dd;",
		Code:        ":memo:",
		Description: "Add or update documentation.",
		Name:        "memo",
	},
	{
		Emoji:       "🚀",
		Entity:      "&#x1f680;",
		Code:        ":rocket:",
		Description: "Deploy stuff.",
		Name:        "rocket",
	},
	{
		Emoji:       "💄",
		Entity:      "&#ff99cc;",
		Code:        ":lipstick:",
		Description: "Add or update the UI and style files.",
		Name:        "lipstick",
	},
	{
		Emoji:       "🎉",
		Entity:      "&#127881;",
		Code:        ":tada:",
		Description: "Begin a project.",
		Name:        "tada",
	},
	{
		Emoji:       "✅",
		Entity:      "&#x2705;",
		Code:        ":white_check_mark:",
		Description: "Add or update tests.",
		Name:        "white-check-mark",
	},
	{
		Emoji:       "🔒",
		Entity:      "&#x1f512;",
		Code:        ":lock:",
		Description: "Fix security issues.",
		Name:        "lock",
	},
	{
		Emoji:       "🔖",
		Entity:      "&#x1f516;",
		Code:        ":bookmark:",
		Description: "Release / Version tags.",
		Name:        "bookmark",
	},
	{
		Emoji:       "🚨",
		Entity:      "&#x1f6a8;",
		Code:        ":rotating_light:",
		Description: "Fix compiler / linter warnings.",
		Name:        "rotating-light",
	},
	{
		Emoji:       "🚧",
		Entity:      "&#x1f6a7;",
		Code:        ":construction:",
		Description: "Work in progress.",
		Name:        "construction",
	},
	{
		Emoji:       "💚",
		Entity:      "&#x1f49a;",
		Code:        ":green_heart:",
		Description: "Fix CI Build.",
		Name:        "green-heart",
	},
	{
		Emoji:       "⬇️",
		Entity:      "⬇️",
		Code:        ":arrow_down:",
		Description: "Downgrade dependencies.",
		Name:        "arrow-down",
	},
	{
		Emoji:       "⬆️",
		Entity:      "⬆️",
		Code:        ":arrow_up:",
		Description: "Upgrade dependencies.",
		Name:        "arrow-up",
	},
	{
		Emoji:       "📌",
		Entity:      "&#x1F4CC;",
		Code:        ":pushpin:",
		Description: "Pin dependencies to specific versions.",
		Name:        "pushpin",
	},
	{
		Emoji:       "👷",
		Entity:      "&#x1f477;",
		Code:        ":construction_worker:",
		Description: "Add or update CI build system.",
		Name:        "construction-worker",
	},
	{
		Emoji:       "📈",
		Entity:      "&#x1F4C8;",
		Code:        ":chart_with_upwards_trend:",
		Description: "Add or update analytics or track code.",
		Name:        "chart-with-upwards-trend",
	},
	{
		Emoji:       "♻️",
		Entity:      "&#x2672;",
		Code:        ":recycle:",
		Description: "Refactor code.",
		Name:        "recycle",
	},
	{
		Emoji:       "➕",
		Entity:      "&#10133;",
		Code:        ":heavy_plus_sign:",
		Description: "Add a dependency.",
		Name:        "heavy-plus-sign",
	},
	{
		Emoji:       "➖",
		Entity:      "&#10134;",
		Code:        ":heavy_minus_sign:",
		Description: "Remove a dependency.",
		Name:        "heavy-minus-sign",
	},
	{
		Emoji:       "🔧",
		Entity:      "&#x1f527;",
		Code:        ":wrench:",
		Description: "Add or update configuration files.",
		Name:        "wrench",
	},
	{
		Emoji:       "🔨",
		Entity:      "&#128296;",
		Code:        ":hammer:",
		Description: "Add or update development scripts.",
		Name:        "hammer",
	},
	{
		Emoji:       "🌐",
		Entity:      "&#127760;",
		Code:        ":globe_with_meridians:",
		Description: "Internationalization and localization.",
		Name:        "globe-with-meridians",
	},
	{
		Emoji:       "✏️",
		Entity:      "&#59161;",
		Code:        ":pencil2:",
		Description: "Fix typos.",
		Name:        "pencil2",
	},
	{
		Emoji:       "💩",
		Entity:      "&#58613;",
		Code:        ":poop:",
		Description: "Write bad code that needs to be improved.",
		Name:        "poop",
	},
	{
		Emoji:       "⏪",
		Entity:      "&#9194;",
		Code:        ":rewind:",
		Description: "Revert changes.",
		Name:        "rewind",
	},
	{
		Emoji:       "🔀",
		Entity:      "&#128256;",
		Code:        ":twisted_rightwards_arrows:",
		Description: "Merge branches.",
		Name:        "twisted-rightwards-arrows",
	},
	{
		Emoji:       "📦",
		Entity:      "&#1F4E6;",
		Code:        ":package:",
		Description: "Add or update compiled files or packages.",
		Name:        "package",
	},
	{
		Emoji:       "👽",
		Entity:      "&#1F47D;",
		Code:        ":alien:",
		Description: "Update code due to external API changes.",
		Name:        "alien",
	},
	{
		Emoji:       "🚚",
		Entity:      "&#1F69A;",
		Code:        ":truck:",
		Description: "Move or rename resources (e.g.: files, paths, routes).",
		Name:        "truck",
	},
	{
		Emoji:       "📄",
		Entity:      "&#1F4C4;",
		Code:        ":page_facing_up:",
		Description: "Add or update license.",
		Name:        "page-facing-up",
	},
	{
		Emoji:       "💥",
		Entity:      "&#x1f4a5;",
		Code:        ":boom:",
		Description: "Introduce breaking changes.",
		Name:        "boom",
	},
	{
		Emoji:       "🍱",
		Entity:      "&#1F371",
		Code:        ":bento:",
		Description: "Add or update assets.",
		Name:        "bento",
	},
	{
		Emoji:       "♿️",
		Entity:      "&#9855;",
		Code:        ":wheelchair:",
		Description: "Improve accessibility.",
		Name:        "wheelchair",
	},
	{
		Emoji:       "💡",
		Entity:      "&#128161;",
		Code:        ":bulb:",
		Description: "Add or update comments in source code.",
		Name:        "bulb",
	},
	{
		Emoji:       "🍻",
		Entity:      "&#x1f37b;",
		Code:        ":beers:",
		Description: "Write code drunkenly.",
		Name:        "beers",
	},
	{
		Emoji:       "💬",
		Entity:      "&#128172;",
		Code:        ":speech_balloon:",
		Description: "Add or update text and literals.",
		Name:        "speech-balloon",
	},
	{
		Emoji:       "🗃",
		Entity:      "&#128451;",
		Code:        ":card_file_box:",
		Description: "Perform database related changes.",
		Name:        "card-file-box",
	},
	{
		Emoji:       "🔊",
		Entity:      "&#128266;",
		Code:        ":loud_sound:",
		Description: "Add or update logs.",
		Name:        "loud-sound",
	},
	{
		Emoji:       "🔇",
		Entity:      "&#128263;",
		Code:        ":mute:",
		Description: "Remove logs.",
		Name:        "mute",
	},
	{
		Emoji:       "👥",
		Entity:      "&#128101;",
		Code:        ":busts_in_silhouette:",
		Description: "Add or update contributor(s).",
		Name:        "busts-in-silhouette",
	},
	{
		Emoji:       "🚸",
		Entity:      "&#128696;",
		Code:        ":children_crossing:",
		Description: "Improve user experience / usability.",
		Name:        "children-crossing",
	},
	{
		Emoji:       "🏗",
		Entity:      "&#1f3d7;",
		Code:        ":building_construction:",
		Description: "Make architectural changes.",
		Name:        "building-construction",
	},
	{
		Emoji:       "📱",
		Entity:      "&#128241;",
		Code:        ":iphone:",
		Description: "Work on responsive design.",
		Name:        "iphone",
	},
	{
		Emoji:       "🤡",
		Entity:      "&#129313;",
		Code:        ":clown_face:",
		Description: "Mock things.",
		Name:        "clown-face",
	},
	{
		Emoji:       "🥚",
		Entity:      "&#129370;",
		Code:        ":egg:",
		Description: "Add or update an easter egg.",
		Name:        "egg",
	},
	{
		Emoji:       "🙈",
		Entity:      "&#8bdfe7;",
		Code:        ":see_no_evil:",
		Description: "Add or update a .gitignore file.",
		Name:        "see-no-evil",
	},
	{
		Emoji:       "📸",
		Entity:      "&#128248;",
		Code:        ":camera_flash:",
		Description: "Add or update snapshots.",
		Name:        "camera-flash",
	},
	{
		Emoji:       "⚗",
		Entity:      "&#128248;",
		Code:        ":alembic:",
		Description: "Perform experiments.",
		Name:        "alembic",
	},
	{
		Emoji:       "🔍",
		Entity:      "&#128269;",
		Code:        ":mag:",
		Description: "Improve SEO.",
		Name:        "mag",
	},
	{
		Emoji:       "🏷️",
		Entity:      "&#127991;",
		Code:        ":label:",
		Description: "Add or update types.",
		Name:        "label",
	},
	{
		Emoji:       "🌱",
		Entity:      "&#127793;",
		Code:        ":seedling:",
		Description: "Add or update seed files.",
		Name:        "seedling",
	},
	{
		Emoji:       "🚩",
		Entity:      "&#x1F6A9;",
		Code:        ":triangular_flag_on_post:",
		Description: "Add, update, or remove feature flags.",
		Name:        "triangular-flag-on-post",
	},
	{
		Emoji:       "🥅",
		Entity:      "&#x1F945;",
		Code:        ":goal_net:",
		Description: "Catch errors.",
		Name:        "goal-net",
	},
	{
		Emoji:       "💫",
		Entity:      "&#x1f4ab;",
		Code:        ":dizzy:",
		Description: "Add or update animations and transitions.",
		Name:        "animation",
	},
	{
		Emoji:       "🗑",
		Entity:      "&#x1F5D1;",
		Code:        ":wastebasket:",
		Description: "Deprecate code that needs to be cleaned up.",
		Name:        "wastebasket",
	},
}

package main

const (
	EMOJI_COIN         = "\U0001FA99"   // (coin)
	EMOJI_SMILE        = "\U0001F642"   // 🙂
	EMOJI_SUNGLASSES   = "\U0001F60E"   // 😎
	EMOJI_WOW          = "\U0001F604"   // 😄
	EMOJI_DONT_KNOW    = "\U0001F937"   // 🤷
	EMOJI_SAD          = "\U0001F63F"   // 😿
	EMOJI_BICEPS       = "\U0001F4AA"   // 💪
	EMOJI_BUTTON_START = "\U000025B6  " // ▶
	EMOJI_BUTTON_END   = "  \U000025C0" // ◀

	BUTTON_TEXT_PRINT_INTRO       = EMOJI_BUTTON_START + "View Introduction" + EMOJI_BUTTON_END
	BUTTON_TEXT_SKIP_INTRO        = EMOJI_BUTTON_START + "Skip Introduction" + EMOJI_BUTTON_END
	BUTTON_TEXT_BALANCE           = EMOJI_BUTTON_START + "Current Balance" + EMOJI_BUTTON_END
	BUTTON_TEXT_USEFUL_ACTIVITIES = EMOJI_BUTTON_START + "Useful Activities" + EMOJI_BUTTON_END
	BUTTON_TEXT_REWARDS           = EMOJI_BUTTON_START + "Rewards" + EMOJI_BUTTON_END
	BUTTON_TEXT_PRINT_MENU        = EMOJI_BUTTON_START + "MAIN MENU" + EMOJI_BUTTON_END

	BUTTON_CODE_PRINT_INTRO       = "print_intro"
	BUTTON_CODE_SKIP_INTRO        = "skip_intro"
	BUTTON_CODE_BALANCE           = "show_balance"
	BUTTON_CODE_USEFUL_ACTIVITIES = "show_useful_activities"
	BUTTON_CODE_REWARDS           = "show_rewards"
	BUTTON_CODE_PRINT_MENU        = "print_menu"

	TOKEN_NAME_IN_OS             = "improve-yourself-bot"
	UPDATE_CONFIG_TIMEOUT        = 60
	MAX_USER_COINS        uint16 = 500
)

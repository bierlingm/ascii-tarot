package main

type Card struct {
	ID      int
	Numeral string
	Name    string
	Art     string
	Desc    []string
	RDesc   []string
}

// Art inspired by guided visualizations:
// - Portal/Christ/Alexander: door rising, flames, sword in right hand, embrace on left
// - King Gathering Army: grove, table, covenant, community descending
// - Fire Horse/Harp: horse through flames, harp completing sword+embrace triad

var Deck = []Card{
	{0, "0", "The Fool", `
       *
      /|\
     / | \
    |  @  |    
       |      
      /|\     
     / | \    
    ~~~~~~~   
`, []string{"beginnings", "new journeys", "trust", "hope", "optimism", "innocence", "spontaneity"}, []string{"naivete", "inexperienced", "easily duped", "risky", "chaos", "bad judgement"}},

	{1, "I", "The Magician", `
      ___
     ( * )
       |
    o--+--o
       |
      / \
   As Above
   So Below
`, []string{"creative power", "vision", "mastery", "energy", "focus", "control", "self-discipline"}, []string{"indecisive", "poor follow-through", "lacking commitment", "misuse of power"}},

	{2, "II", "The High Priestess", `
    B     J
    |     |
   [=======]
      / \
     / O \
    /     \
   |  ~~~  |
   | TORA  |
`, []string{"intuition", "instinct", "dreams", "mysteries", "deep emotions", "greater powers"}, []string{"repressed intuition", "blocks", "false allies", "hidden enemies"}},

	{3, "III", "The Empress", `
     \*/
    __|__
   /     \
  | Venus |
   \_____/
    |   |
   /|   |\
  ~~~~~~~~
`, []string{"fruitfulness", "fertility", "fulfillment", "sensuality", "beauty", "nurturing", "motherhood"}, []string{"energy wasted", "abilities obstructed", "laziness", "poverty"}},

	{4, "IV", "The Emperor", `
     ___
    /   \
   | .-. |
    \___/
   __|=|__
  /   |   \
 /   /|\   \
 =========
`, []string{"power", "authority", "assertion", "control", "order", "structure", "logic"}, []string{"controlling", "inflexibility", "abuse of power", "coldness"}},

	{5, "V", "The Hierophant", `
    .---.
   ( ~~~ )
    \___/
   __|^|__
  / B | B \
 |  --|--  |
  \  / \  /
   ======
`, []string{"spirituality", "education", "higher learning", "mentor", "seeking guidance"}, []string{"breaking rules", "anarchy", "rebellion", "outmoded beliefs"}},

	// THE LOVERS - inspired by sword (boundary) + embrace (love) imagery
	{6, "VI", "The Lovers", `
       *
      ~~~
   o     o
  /|\   /|\
   |--<3--|
  / \   / \
 SWORD  EMBRACE
`, []string{"love", "passion", "relationships", "shared values", "balance", "choices"}, []string{"infidelity", "imbalance", "jealousy", "arguments", "broken hearts"}},

	// THE CHARIOT - inspired by fire horse visualization
	{7, "VII", "The Chariot", `
      ___
     [_*_]
       |
  __/=====\__
 / ~~ ### ~~ \
|  *   |   *  |
 \____/_\____/
  oo       oo
`, []string{"energy", "drive", "ambition", "moving forward", "victory", "triumph"}, []string{"unfocused", "energy scattered", "negativity", "lacking self-control"}},

	// STRENGTH - the embrace, calm authority
	{8, "VIII", "Strength", `
      o
     /|\
    ~~|~~
   /  |  \
  | ____ |
   / oo \
  |  <>  |
   \____/
`, []string{"strength from within", "faith", "bravery", "patience", "self-control"}, []string{"weakness", "lack of self-control", "lack of courage", "self doubt"}},

	// THE HERMIT - introspection, the king alone before gathering
	{9, "IX", "The Hermit", `
      *
     ( )
      |
     /|\
    / | \
   |  |  |
      |
     / \
    ~~~~
`, []string{"introspection", "meditation", "solitary endeavors", "wisdom", "study"}, []string{"isolation", "loneliness", "fear of being alone", "misfit"}},

	// WHEEL OF FORTUNE - portal rising, threshold
	{10, "X", "Wheel of Fortune", `
    .---.
   /     \
  | T A R |
  |O  *  O|
  | A R T |
   \     /
    '---'
`, []string{"good luck", "destiny", "karma", "serendipity", "chance meetings"}, []string{"bad luck", "bad news", "problems", "delays", "unexpected changes"}},

	// JUSTICE - Alexander's sword, discernment
	{11, "XI", "Justice", `
     ___
    /   \
   | === |
    \___/
   /  |  \
  |   |   |
 ===  |  ===
      |
     / \
`, []string{"equilibrium", "justice", "fairness", "honor", "resolution", "legal dealings"}, []string{"disequilibrium", "injustice", "delays", "dishonesty"}},

	// HANGED MAN - suspension, letting go
	{12, "XII", "The Hanged Man", `
   ========
      ||
      ||
      /\
     /  \
    / o  \
   |  |   |
   | / \  |
    \___/
`, []string{"sacrifice", "understanding", "letting go", "detachment", "suspension"}, []string{"selfishness", "pain", "cannot let go", "materialism"}},

	// DEATH - transformation, walking through portal
	{13, "XIII", "Death", `
     ___
    / X \
   |X   X|
    \___/
      |
   /=====\
  |  ~~~  |
   \_____/
    TRANS
`, []string{"change", "transformation", "transition", "end", "new beginning", "renewal"}, []string{"boredom", "sloth", "depression", "unable to move on"}},

	// TEMPERANCE - the harp, harmonizing multiple strings
	{14, "XIV", "Temperance", `
      o
     /|\
    / | \
   |  *  |
  (  |||  )
   \ ||| /
    \||//
     \/
`, []string{"moderation", "balance", "harmony", "resolution", "partnership"}, []string{"discordance", "disagreements", "outbursts", "out of balance"}},

	// THE DEVIL - the temptation before the portal
	{15, "XV", "The Devil", `
    /\ /\
   (  V  )
    \___/
   __|^|__
  /       \
 | o  *  o |
  \_______/
   BONDAGE
`, []string{"bondage", "obsession", "materialism", "imbalanced relationships", "lust"}, []string{"chains broken", "severing ties", "exhaustion", "shifting values"}},

	// THE TOWER - disruptive change, old structures falling
	{16, "XVI", "The Tower", `
    /\
   /##\
  |####|
  |## *|
  | ## |
  |*  #|
 /|    |\
  ~~~~~~
`, []string{"disruptive change", "extreme changes", "anger", "disaster", "ruin"}, []string{"upheaval", "imprisonment", "freedom for a price", "volatility"}},

	// THE STAR - hope, the gleaming spiral path
	{17, "XVII", "The Star", `
      *
     /|\
    * | *
      |
      o
     /|\
    / | \
   ~~ | ~~
`, []string{"hope", "goals within reach", "inspiration", "healing", "new approach"}, []string{"pessimism", "faithlessness", "hopelessness", "self-doubt"}},

	// THE MOON - the threshold between realms
	{18, "XVIII", "The Moon", `
    .--.
   (    )
    '--'
      |
   /     \
  | o   o |
   \_____/
    ~~~~~
`, []string{"concealed changes", "dreams", "intuition", "portent", "secrets", "trickery"}, []string{"insomnia", "unhappiness", "irrational action", "paranoid"}},

	// THE SUN - descending to bless the community
	{19, "XIX", "The Sun", `
   \  |  /
    \ | /
  ---(**)---
    / | \
   /  |  \
      o
     /|\
    / | \
`, []string{"positivity", "happiness", "vitality", "warmth", "success", "vacation"}, []string{"half success", "temporary difficulties", "delayed success"}},

	// JUDGEMENT - the covenant, awakening
	{20, "XX", "Judgement", `
     .-.
    (   )
     '-'
   /  |  \
  |   |   |
  o   |   o
 /|\  |  /|\
    ====
`, []string{"rebirth", "awakening", "fresh start", "renewal", "forgiveness"}, []string{"uninspired", "boredom", "refusal to self-assess", "fear of change"}},

	// THE WORLD - completion, the fire horse arrives
	{21, "XXI", "The World", `
   .-----.
  /   *   \
 |  \o/    |
 | --|--   |
 |  / \    |
  \_______/
  COMPLETE
`, []string{"completion", "conclusion", "success", "next level", "new era", "recognition"}, []string{"lesson not learned", "completion delayed", "stagnation"}},

	// Minor Arcana - Wands (keeping simpler)
	{22, "Ace", "Ace of Wands", `
      *
      |
     /|\
      |
      |
     ~~~
`, []string{"creation", "origin", "venture", "invention", "beginning"}, []string{"fall", "delays", "ruin", "spiritual blocks"}},
	{23, "II", "Two of Wands", "", []string{"magnificence", "dominion", "productive relationships", "work friendships"}, []string{"trouble", "fears", "hollow success", "miscommunication"}},
	{24, "III", "Three of Wands", "", []string{"manifesting dreams", "fruitful effort", "success", "trade", "discovery"}, []string{"halt in difficulties", "unrealistic expectations"}},
	{25, "IV", "Four of Wands", "", []string{"new home", "rest", "refuge", "harvest", "harmony", "prosperity"}, []string{"growth", "prosperity", "happiness", "beauty"}},
	{26, "V", "Five of Wands", "", []string{"arguments", "conflicts", "testing circumstances", "power struggles"}, []string{"litigation", "court battles", "inner conflict", "anger"}},
	{27, "VI", "Six of Wands", "", []string{"victory", "hard won battles", "legal triumphs", "good news", "success"}, []string{"enemy at gate", "troubles caused by others", "lack of reward"}},
	{28, "VII", "Seven of Wands", "", []string{"standing up for yourself", "tenacity", "struggle", "faith in success"}, []string{"fear of failure", "insecurity", "risk", "anxiety"}},
	{29, "VIII", "Eight of Wands", "", []string{"swiftness", "haste", "action", "travel", "change in career"}, []string{"delays", "dispute", "hurrying creates mistakes"}},
	{30, "IX", "Nine of Wands", "", []string{"strength", "resilience", "one last challenge", "persistence"}, []string{"obstacles", "insecurities", "loss of will", "stress"}},
	{31, "X", "Ten of Wands", "", []string{"oppression", "hard work", "ambition", "burden", "drive"}, []string{"schemes", "difficulties", "biting off more than you can chew"}},
	{32, "Page", "Page of Wands", "", []string{"child", "creative ideas", "education", "enthusiasm"}, []string{"bad news", "setbacks", "announcements"}},
	{33, "Knight", "Knight of Wands", "", []string{"young man", "charming", "bright", "creative", "adventure"}, []string{"con artist", "fast-talker", "selfish", "rogue"}},
	{34, "Queen", "Queen of Wands", "", []string{"exuberant woman", "enthusiastic", "independent", "passionate"}, []string{"jealousy", "vengeful", "drama queen", "infidelity"}},
	{35, "King", "King of Wands", "", []string{"honest man", "fair", "natural leader", "successful", "humorous"}, []string{"sober", "dogmatic", "violent temper", "impulsive"}},
	// Cups
	{36, "Ace", "Ace of Cups", `
    .---.
   /     \
  |   <3  |
   \_____/
     | |
    _| |_
`, []string{"heart's home", "abundance", "love", "passion", "growth"}, []string{"chaos", "need for replenishment", "one-sided relationships"}},
	{37, "II", "Two of Cups", "", []string{"union", "love", "connections", "harmony", "romantic attachments"}, []string{"emotional battles", "misunderstandings", "separations"}},
	{38, "III", "Three of Cups", "", []string{"celebration", "parties", "social events", "weddings", "fulfillment"}, []string{"excess", "lack of sleep", "hedonism", "overindulgence"}},
	{39, "IV", "Four of Cups", "", []string{"emotional boredom", "depression", "stuck in a rut", "fear of love"}, []string{"fear of being alone", "new challenges", "new hobby"}},
	{40, "V", "Five of Cups", "", []string{"loss", "destruction of ideals", "broken dreams", "old wounds"}, []string{"reserved hope", "new chances", "rebuilding relationships"}},
	{41, "VI", "Six of Cups", "", []string{"happy memories", "nostalgia", "old relationships", "rewards for past"}, []string{"clinging to the past", "blaming the past", "self-doubt"}},
	{42, "VII", "Seven of Cups", "", []string{"illusions", "wishful thinking", "imagination", "temporary success"}, []string{"desire", "deceptions", "seeing only what you want"}},
	{43, "VIII", "Eight of Cups", "", []string{"abandoning relationships", "moving away emotionally", "travel"}, []string{"hopelessness", "depression", "search for meaning"}},
	{44, "IX", "Nine of Cups", "", []string{"joy", "happiness", "fulfillment", "well-being", "living with ease"}, []string{"greedy", "smug", "taking things for granted"}},
	{45, "X", "Ten of Cups", "", []string{"joy in abundance", "lasting happiness", "happy homes", "dreams come true"}, []string{"broken home", "quarrels", "discord", "loss of friendship"}},
	{46, "Page", "Page of Cups", "", []string{"thoughtful", "imaginative", "meditative", "loving", "dreams"}, []string{"emotionally immature", "seduction", "illusion"}},
	{47, "Knight", "Knight of Cups", "", []string{"romantic dreamer", "spiritual", "creative", "playful", "idealistic"}, []string{"talented but unrealistic", "lazy", "living in a daze"}},
	{48, "Queen", "Queen of Cups", "", []string{"devoted woman", "fantasy", "profound intuition", "artistic"}, []string{"perverse", "overly emotional", "manipulative"}},
	{49, "King", "King of Cups", "", []string{"emotionally balanced", "quiet confidence", "creative", "intuitive"}, []string{"intense", "two-faced", "unfaithful", "addiction"}},
	// Swords
	{50, "Ace", "Ace of Swords", `
      *
     /|\
    / | \
   |  |  |
      |
     / \
    =====
`, []string{"new thoughts", "new possibilities", "triumph", "power", "conquest"}, []string{"sudden changes", "imbalanced mind", "quarrels"}},
	{51, "II", "Two of Swords", "", []string{"friendship", "agreement", "balance", "stalemate", "justice"}, []string{"caution", "confusion", "lies and deceit"}},
	{52, "III", "Three of Swords", "", []string{"removal", "betrayal", "absence", "sorrow", "heartbreak"}, []string{"emotional stress", "anxiety", "error", "distraction"}},
	{53, "IV", "Four of Swords", "", []string{"rest", "relaxation", "retreat", "need for peace", "solitude"}, []string{"restlessness", "health issues", "limited progress"}},
	{54, "V", "Five of Swords", "", []string{"loss", "destruction", "failure", "damage", "sense of unfairness"}, []string{"releasing pride", "defeat", "acceptance", "burial"}},
	{55, "VI", "Six of Swords", "", []string{"leaving strife behind", "pleasant times", "alleviating stress", "travel"}, []string{"struggle", "confession", "delayed victory"}},
	{56, "VII", "Seven of Swords", "", []string{"sporadic effort", "huge volume of ideas", "schemes", "idea theft"}, []string{"confidence", "hype", "false promises", "manipulation"}},
	{57, "VIII", "Eight of Swords", "", []string{"bondage", "stuck", "crisis", "restriction", "feeling trapped"}, []string{"release", "unsettled", "open to new perspectives"}},
	{58, "IX", "Nine of Swords", "", []string{"oppression", "anxiety", "stress", "worries", "sleepless nights"}, []string{"imprisonment", "torment", "dark night of the soul"}},
	{59, "X", "Ten of Swords", "", []string{"betrayal", "backstabbing", "slander", "gossip", "rock bottom"}, []string{"depression", "negative thinking", "pessimism"}},
	{60, "Page", "Page of Swords", "", []string{"brilliance", "curiosity", "strong ideas", "perceptive"}, []string{"meddling", "vindictive", "mean", "aggression"}},
	{61, "Knight", "Knight of Swords", "", []string{"sharp", "egoist", "bravery", "competitive", "fearless"}, []string{"recklessness", "rival", "tyranny", "collapse"}},
	{62, "Queen", "Queen of Swords", "", []string{"witty", "sharp mind", "sarcastic", "freedom", "logical"}, []string{"evil-minded", "predator", "unfeeling", "spiteful"}},
	{63, "King", "King of Swords", "", []string{"intellectual", "powerful authority", "law", "achiever"}, []string{"authoritarian", "rude", "cruel", "strong enemy"}},
	// Pentacles
	{64, "Ace", "Ace of Pentacles", `
    .---.
   /  $  \
  |   *   |
   \_____/
    ~~~~~
`, []string{"contentment", "increase", "new sources of money", "prosperity"}, []string{"greed", "fear of scarcity", "hoarding"}},
	{65, "II", "Two of Pentacles", "", []string{"financial balance", "extra work", "new skills", "entrepreneur"}, []string{"financial discord", "lack of focus", "growing debts"}},
	{66, "III", "Three of Pentacles", "", []string{"one's forte", "apprenticeship", "learning new skill", "self-improvement"}, []string{"mediocrity", "unwilling to invest time", "holding yourself back"}},
	{67, "IV", "Four of Pentacles", "", []string{"financial stability", "healthy income", "modest gains", "inheritance"}, []string{"social climber", "greed", "miserly", "materialism"}},
	{68, "V", "Five of Pentacles", "", []string{"despair", "financial hardships", "lack of faith", "troublesome times"}, []string{"discord", "wantonness", "chaos", "ruin"}},
	{69, "VI", "Six of Pentacles", "", []string{"gifts", "help given and received", "philanthropy", "charity"}, []string{"poor spending habits", "gambling losses", "bribes", "theft"}},
	{70, "VII", "Seven of Pentacles", "", []string{"long-term investment", "vision", "growth of business", "perseverance"}, []string{"mental exhaustion", "anxiety about money", "limited reward"}},
	{71, "VIII", "Eight of Pentacles", "", []string{"new skills", "apprenticeship", "education", "expansion"}, []string{"emptied ambitions", "limitations at work", "no new potential"}},
	{72, "IX", "Nine of Pentacles", "", []string{"prosperity", "money flowing", "the finer things", "accomplishment"}, []string{"lack of funds", "bad investments", "cancelled projects"}},
	{73, "X", "Ten of Pentacles", "", []string{"gains", "family affairs", "gifts", "inheritance", "tax refunds"}, []string{"financial dreams", "unexpected expenses", "burdeonsome family"}},
	{74, "Page", "Page of Pentacles", "", []string{"studiousness", "good news about money", "a raise", "profits"}, []string{"lacking progress", "introverted", "illogical", "moody"}},
	{75, "Knight", "Knight of Pentacles", "", []string{"useful", "determined", "ambitious", "successful", "loyal"}, []string{"apathy", "stagnation", "laziness", "careless with money"}},
	{76, "Queen", "Queen of Pentacles", "", []string{"lavish", "generous", "luxury", "warm-hearted", "maternal"}, []string{"greed", "gold-digger", "overly ambitious", "untrustworthy"}},
	{77, "King", "King of Pentacles", "", []string{"wealthy", "valor", "skilled", "responsibility", "steadfast"}, []string{"miser", "corruption", "bully", "jealousy"}},
}

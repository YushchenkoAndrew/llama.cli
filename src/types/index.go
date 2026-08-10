package types

// Models represents the outer JSON wrapper object.
type Models struct {
	Data []Model `json:"data"`
}

// Model represents the inner details of the model payload.
type Model struct {
	ID      string   `json:"id"`
	Aliases []string `json:"aliases"`
	Tags    []string `json:"tags"` // Using string slice assuming tags are strings
	Object  string   `json:"object"`
	OwnedBy string   `json:"owned_by"`
	Created int64    `json:"created"` // int64 is ideal for Unix timestamps
	Status  struct {
		Value  string   `json:"value"`
		Args   []string `json:"args"`
		Preset string   `json:"preset"`
	} `json:"status"` // Status handles the current running or unloaded state of the model.
	Architecture struct {
		InputModalities  []string `json:"input_modalities"`
		OutputModalities []string `json:"output_modalities"`
	} `json:"architecture"` // Architecture details the model's supported modalities.
	Source    string     `json:"source"`
	CanRemove bool       `json:"can_remove"`
	Meta      *ModelMeta `json:"meta,omitempty"`
}

type ModelMeta struct {
	VocabType int    `json:"vocab_type"`
	NVocab    int    `json:"n_vocab"`
	NCtx      int    `json:"n_ctx"`
	NCtxTrain int    `json:"n_ctx_train"`
	NEmbd     int    `json:"n_embd"`
	NParams   int64  `json:"n_params"` // int64 prevents overflow for large parameters (7B+)
	Size      int64  `json:"size"`     // int64 prevents overflow for large file sizes in bytes
	Ftype     string `json:"ftype"`
}

type Props struct {
	DefaultGenerationSettings struct {
		Params struct {
			Seed               uint32   `json:"seed"` // Fits 4294967295 perfectly
			Temperature        float64  `json:"temperature"`
			DynatempRange      float64  `json:"dynatemp_range"`
			DynatempExponent   float64  `json:"dynatemp_exponent"`
			TopK               int      `json:"top_k"`
			TopP               float64  `json:"top_p"`
			MinP               float64  `json:"min_p"`
			TopNSigma          float64  `json:"top_n_sigma"`
			XtcProbability     float64  `json:"xtc_probability"`
			XtcThreshold       float64  `json:"xtc_threshold"`
			TypicalP           float64  `json:"typical_p"`
			RepeatLastN        int      `json:"repeat_last_n"`
			RepeatPenalty      float64  `json:"repeat_penalty"`
			PresencePenalty    float64  `json:"presence_penalty"`
			FrequencyPenalty   float64  `json:"frequency_penalty"`
			DryMultiplier      float64  `json:"dry_multiplier"`
			DryBase            float64  `json:"dry_base"`
			DryAllowedLength   int      `json:"dry_allowed_length"`
			DryPenaltyLastN    int      `json:"dry_penalty_last_n"`
			Mirostat           int      `json:"mirostat"`
			MirostatTau        float64  `json:"mirostat_tau"`
			MirostatEta        float64  `json:"mirostat_eta"`
			AdaptiveTarget     float64  `json:"adaptive_target"`
			AdaptiveDecay      float64  `json:"adaptive_decay"`
			MaxTokens          int      `json:"max_tokens"`
			NPredict           int      `json:"n_predict"`
			NKeep              int      `json:"n_keep"`
			NDiscard           int      `json:"n_discard"`
			IgnoreEos          bool     `json:"ignore_eos"`
			Stream             bool     `json:"stream"`
			NProbs             int      `json:"n_probs"`
			MinKeep            int      `json:"min_keep"`
			ChatFormat         string   `json:"chat_format"`
			ReasoningFormat    string   `json:"reasoning_format"`
			ReasoningInContent bool     `json:"reasoning_in_content"`
			GenerationPrompt   string   `json:"generation_prompt"`
			Samplers           []string `json:"samplers"`
			SpeculativeTypes   string   `json:"speculative.types"` // Handles key with a period
			TimingsPerToken    bool     `json:"timings_per_token"`
			PostSamplingProbs  bool     `json:"post_sampling_probs"`
			BackendSampling    bool     `json:"backend_sampling"`
			Lora               []any    `json:"lora"`
		} `json:"params"`
		NCtx int `json:"n_ctx"`
	} `json:"default_generation_settings"`
	TotalSlots int    `json:"total_slots"`
	ModelAlias string `json:"model_alias"`
	ModelFtype string `json:"model_ftype"`
	ModelPath  string `json:"model_path"`
	Modalities struct {
		Vision bool `json:"vision"`
		Video  bool `json:"video"`
		Audio  bool `json:"audio"`
	} `json:"modalities"`
	MediaMarker      string         `json:"media_marker"`
	EndpointSlots    bool           `json:"endpoint_slots"`
	EndpointProps    bool           `json:"endpoint_props"`
	EndpointMetrics  bool           `json:"endpoint_metrics"`
	Ui               bool           `json:"ui"`
	UiSettings       map[string]any `json:"ui_settings"` // Flexible map for variable UI configs
	ChatTemplate     string         `json:"chat_template"`
	ChatTemplateCaps struct {
		SupportsObjectArguments   bool `json:"supports_object_arguments"`
		SupportsParallelToolCalls bool `json:"supports_parallel_tool_calls"`
		SupportsPreserveReasoning bool `json:"supports_preserve_reasoning"`
		SupportsStringContent     bool `json:"supports_string_content"`
		SupportsSystemRole        bool `json:"supports_system_role"`
		SupportsToolCalls         bool `json:"supports_tool_calls"`
		SupportsTools             bool `json:"supports_tools"`
		SupportsTypedContent      bool `json:"supports_typed_content"`
	} `json:"chat_template_caps"`
	BosToken         string `json:"bos_token"`
	EosToken         string `json:"eos_token"`
	BuildInfo        string `json:"build_info"`
	IsSleeping       bool   `json:"is_sleeping"`
	CorsProxyEnabled bool   `json:"cors_proxy_enabled"`
}

type Slot struct {
	ID                     int  `json:"id"`
	NCtx                   int  `json:"n_ctx"`
	Speculative            bool `json:"speculative"`
	IsProcessing           bool `json:"is_processing"`
	IDTask                 int  `json:"id_task"`
	NPromptTokens          int  `json:"n_prompt_tokens"`
	NPromptTokensProcessed int  `json:"n_prompt_tokens_processed"`
	NPromptTokensCache     int  `json:"n_prompt_tokens_cache"`
	Params                 struct {
		Seed               uint32   `json:"seed"`
		Temperature        float64  `json:"temperature"`
		DynatempRange      float64  `json:"dynatemp_range"`
		DynatempExponent   float64  `json:"dynatemp_exponent"`
		TopK               int      `json:"top_k"`
		TopP               float64  `json:"top_p"`
		MinP               float64  `json:"min_p"`
		TopNSigma          float64  `json:"top_n_sigma"`
		XtcProbability     float64  `json:"xtc_probability"`
		XtcThreshold       float64  `json:"xtc_threshold"`
		TypicalP           float64  `json:"typical_p"`
		RepeatLastN        int      `json:"repeat_last_n"`
		RepeatPenalty      float64  `json:"repeat_penalty"`
		PresencePenalty    float64  `json:"presence_penalty"`
		FrequencyPenalty   float64  `json:"frequency_penalty"`
		DryMultiplier      float64  `json:"dry_multiplier"`
		DryBase            float64  `json:"dry_base"`
		DryAllowedLength   int      `json:"dry_allowed_length"`
		DryPenaltyLastN    int      `json:"dry_penalty_last_n"`
		Mirostat           int      `json:"mirostat"`
		MirostatTau        float64  `json:"mirostat_tau"`
		MirostatEta        float64  `json:"mirostat_eta"`
		AdaptiveTarget     float64  `json:"adaptive_target"`
		AdaptiveDecay      float64  `json:"adaptive_decay"`
		MaxTokens          int      `json:"max_tokens"`
		NPredict           int      `json:"n_predict"`
		NKeep              int      `json:"n_keep"`
		NDiscard           int      `json:"n_discard"`
		IgnoreEos          bool     `json:"ignore_eos"`
		Stream             bool     `json:"stream"`
		NProbs             int      `json:"n_probs"`
		MinKeep            int      `json:"min_keep"`
		ChatFormat         string   `json:"chat_format"`
		ReasoningFormat    string   `json:"reasoning_format"`
		ReasoningInContent bool     `json:"reasoning_in_content"`
		GenerationPrompt   string   `json:"generation_prompt"`
		Samplers           []string `json:"samplers"`
		SpeculativeTypes   string   `json:"speculative.types"`
		TimingsPerToken    bool     `json:"timings_per_token"`
		PostSamplingProbs  bool     `json:"post_sampling_probs"`
		BackendSampling    bool     `json:"backend_sampling"`
		Lora               []any    `json:"lora"`
	} `json:"params"`
	NextToken []struct {
		HasNextToken bool `json:"has_next_token"`
		HasNewLine   bool `json:"has_new_line"`
		NRemain      int  `json:"n_remain"`
		NDecoded     int  `json:"n_decoded"`
	} `json:"next_token"`
}

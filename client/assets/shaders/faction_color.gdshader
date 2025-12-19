shader_type canvas_item;

// Faction color shader for ships
// Applies faction color to grayscale base texture
// White areas become faction color, shadows are preserved

uniform vec4 faction_color : source_color = vec4(1.0, 1.0, 1.0, 1.0);
uniform float color_intensity : hint_range(0.0, 1.0) = 0.8;

void fragment() {
	vec4 tex = texture(TEXTURE, UV);
	
	// Calculate luminance (grayscale value)
	float luminance = dot(tex.rgb, vec3(0.299, 0.587, 0.114));
	
	// Mix faction color with luminance to preserve shading
	vec3 colored = faction_color.rgb * luminance;
	vec3 final_color = mix(tex.rgb, colored, color_intensity);
	
	COLOR = vec4(final_color, tex.a);
}

#include <stdlib.h>
#include <glib.h>
#include <glib-object.h>
#include <gio/gio.h>

// This exists because cGo doesn't support variadic functions
void
g_variant_builder_add_variant(
    GVariantBuilder *builder,
    const gchar *key,
    GVariant *value
) {
    g_variant_builder_add(builder, "{s@v}", key, value);
}
